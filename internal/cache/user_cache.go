package cache

import (
	"context"
	"encoding/json"
	"myproject/internal/models"
	"time"

	"github.com/go-redis/redis"
)

func (c *Cache) AddTokenToCache(ctx context.Context, id string, token string) error {
	err := c.rdb.Set(ctx, id, token, time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) GetTokenFromCache(ctx context.Context, id string) (string, error) {
	val, err := c.rdb.Get(ctx, id).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (c *Cache) AddAllUsersToCache(ctx context.Context, users []models.FetchAllUsers) error {
	usersJSON, err := json.Marshal(users)
	if err != nil {
		return err
	}
	err = c.rdb.Set(ctx, "all_users", usersJSON, time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) GetAllUsersFromCache(ctx context.Context, key string) ([]models.FetchAllUsers, error) {
	usersJSON, err := c.rdb.Get(ctx, "all_users").Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var users []models.FetchAllUsers
	err = json.Unmarshal([]byte(usersJSON), &users) 
	if err != nil {
		return nil, err
	}

	return users, nil
}
