package main

import (
	"log"
	"go-idempotent-payments/internal/payments"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	
	h := payments.NewHandler(rdb)

	r := gin.Default()
	r.POST("/pay", h.Process)
	r.GET("/health", func(c *gin.Context){ c.JSON(200, gin.H{"status":"ok"}) })
	
	log.Println("API :8080 with Redis idempotency")
	r.Run(":8080")
}
