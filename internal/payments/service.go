package payments

import (
	"context"
	"time"
	"github.com/go-redis/redis/v8"
)

type Service struct {
	redis *redis.Client
}

func NewService(rdb *redis.Client) *Service { return &Service{redis: rdb} }

func (s *Service) ProcessPayment(ctx context.Context, key string, amount int64) (string, bool, error) {
	// Check idempotency
	val, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		return val, true, nil // replay
	}

	result := "pay_" + time.Now().Format("20060102150405")
	s.redis.Set(ctx, key, result, 24*time.Hour)
	return result, false, nil
}
