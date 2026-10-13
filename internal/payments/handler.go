package payments

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Handler struct { service *Service }

func NewHandler(rdb *redis.Client) *Handler {
	return &Handler{service: NewService(rdb)}
}

func (h *Handler) Process(c *gin.Context) {
	key := c.GetHeader("Idempotency-Key")
	if key == "" {
		c.JSON(400, gin.H{"error":"Idempotency-Key required"}); return
	}
	var req struct{ Amount int64 `json:"amount" binding:"required"` }
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error":err.Error()}); return
	}

	result, isReplay, _ := h.service.ProcessPayment(c.Request.Context(), key, req.Amount)
	
	status := "processed"
	if isReplay { status = "idempotent_replay" }
	
	c.JSON(200, gin.H{"status": status, "payment_id": result})
}
