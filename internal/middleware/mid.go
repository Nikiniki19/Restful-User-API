package middleware

import (
	"myproject/internal/auth"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	auth auth.AuthInterface
}
type MiddlewareInterface interface {
	Authenticate() gin.HandlerFunc
	RoleAuthMiddleware(requiredRoles ...string) gin.HandlerFunc
}

func NewMiddleWare(auth auth.AuthInterface) (MiddlewareInterface, error) {
	return &Middleware{
		auth: auth,
	}, nil
}
