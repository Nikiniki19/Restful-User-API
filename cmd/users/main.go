package main

import (
	"myproject/internal/auth"
	"myproject/internal/cache"
	"myproject/internal/database"
	"myproject/internal/handler"
	"myproject/internal/middleware"
	"myproject/internal/repository"
	"myproject/internal/server"
	"myproject/internal/service"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Connected to the database")
	dbConn, err := database.ConnectToDatabase()
	if err != nil {
		return
	}
	repo, err := repository.NewRepository(dbConn)
	if err != nil {
		return
	}
	rdb, err := database.RedisConnect()
	if err != nil {
		return
	}
	cache := cache.NewCache(*rdb)

	auth, err := auth.NewAuth(auth.SignInKey)
	if err != nil {
		return
	}
	mid, err := middleware.NewMiddleWare(auth)
	if err != nil {
		return
	}
	service, err := service.NewService(repo, auth, cache)
	if err != nil {
		return
	}
	handler, err := handler.NewHandler(service)
	if err != nil {
		return
	}

	log.Info().Msg("Starting the application")
	server.StartApplicationServer(handler, mid)

}
