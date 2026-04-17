package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func correlationIDMiddleware(c *gin.Context) {
	correlationID := c.GetHeader("X-Correlation-ID")
	if correlationID == "" {
		correlationID = uuid.New().String()
	}
	c.Set("CorrelationID", correlationID)
	c.Header("X-Correlation-ID", correlationID)
	c.Next()
}
