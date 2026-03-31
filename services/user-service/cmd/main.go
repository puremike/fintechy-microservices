package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	utils "github.com/puremike/fintechy-microservices/shared/utils"

	grpcserver "google.golang.org/grpc"
)

var (
	port = utils.GetEnvString("PORT", "8100")
	// dbAddr = utils.GetEnvString("DB_ADDR", "postgres://admin:finuserdb@localhost:5433/finuserdb?sslmode=disable")
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// db, err := db.NewUserDB(dbAddr)
	// if err != nil {
	// 	utils.Logger().Fatalw("Failed to connect to user database:", "error", err)
	// }
	// defer db.Close()

	// utils.Logger().Infow("Successfully connected to user database")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-shutdown
		utils.Logger().Infow("Received user service shutdown signal:", "signal", sig.String())
		cancel()
	}()

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		utils.Logger().Infow("User service failed to listen on port:", "port", port, "error", err)
	}

	gRPCServer := grpcserver.NewServer()
	// GRPCHANDLER

	go func() {
		utils.Logger().Infow("User service server listens on port:", "port", port)

		if err := gRPCServer.Serve(listen); err != nil {
			utils.Logger().Fatalw("User service failed to serve on:", "port", listen.Addr().String(), "error", err)
		}
	}()

	<-ctx.Done() // wait for shutdown
	utils.Logger().Infow("User service is shutting down...")
	gRPCServer.GracefulStop()
	utils.Logger().Infow("User service has shut down gracefully.")
}
