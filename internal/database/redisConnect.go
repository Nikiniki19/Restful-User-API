package database

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func RedisConnect() (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Error().Err(err).Msg("Connection to the redis is closed")
		return nil, err
	}
	return rdb, nil
}
