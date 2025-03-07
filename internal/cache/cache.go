package cache

import (
	"context"
	"myproject/internal/models"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	rdb redis.Client
}

type CacheInterface interface {
	AddTokenToCache(ctx context.Context, id string, token string) error
	GetTokenFromCache(ctx context.Context, id string) (string, error)
	AddAllUsersToCache(ctx context.Context, users []models.FetchAllUsers) error
	GetAllUsersFromCache(ctx context.Context, key string) ([]models.FetchAllUsers, error)
}

func NewCache(rdb redis.Client) CacheInterface {
	return &Cache{
		rdb: rdb,
	}
}
