package handler

import (
	"myproject/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Ser service.ServiceInterface
}

type HandlerInterface interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
	FetchAllUsers(c *gin.Context)
	FetchUserByID(c *gin.Context)
	UpdateUserByID(c *gin.Context)
	DeleteUserByID(c *gin.Context)
}

func NewHandler(ser service.ServiceInterface) (HandlerInterface, error) {
	return &Handler{
		Ser: ser,
	}, nil
}
