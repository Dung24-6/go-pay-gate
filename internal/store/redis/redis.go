package redis

import (
	"context"
	"fmt"

	"github.com/Dung24-6/go-pay-gate/internal/config"
	goredis "github.com/redis/go-redis/v9" // Rename import để tránh conflict
)

// Client wraps redis client
type Client struct {
	*goredis.Client
}

// NewRedisClient creates a new Redis client
func NewRedisClient(cfg *config.RedisConfig) (*Client, error) {
	client := goredis.NewClient(&goredis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
		MaxRetries:   cfg.MaxRetries,
	})

	// Test connection
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect redis: %w", err)
	}

	return &Client{client}, nil
}
