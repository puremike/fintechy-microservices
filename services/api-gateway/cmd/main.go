package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/puremike/fintechy-microservices/shared/utils"
)

var (
	port = utils.GetEnvString("PORT", "8070")
)

func main() {

	mux := route()
	if err := server(mux); err != nil {
		utils.Logger().Fatalw("API Gateway failed to start:", "error", err)
	}
}

func server(mux http.Handler) error {

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute * 1,
	}

	serverError := make(chan error, 1)

	go func() {
		utils.Logger().Infow("API Gateway is running on port:", "port", port)
		serverError <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverError:
		utils.Logger().Errorw("API Gateway encountered an error:", "error", err)

	case sig := <-shutdown:
		utils.Logger().Infow("API Gateway received a shutdown signal:", "signal", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		utils.Logger().Infow("API Gateway is shutting down...")

		if err := server.Shutdown(ctx); err != nil {
			utils.Logger().Errorw("API Gateway shutdown error:", "error", err)

			server.Close()
		}
	}

	return nil
}

func route() http.Handler {

	g := gin.Default()

	g.GET("/check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "API Gateway is healthy"})
	})

	time.Sleep(time.Second * 7)

	return g
}
