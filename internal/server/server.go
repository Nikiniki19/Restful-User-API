package server

import (
	"myproject/internal/handler"
	"myproject/internal/middleware"

	"github.com/gin-gonic/gin"
)

func StartApplicationServer(h *handler.Handler, m *middleware.Middleware) {
	router := gin.Default()

	// Public routes (No authentication required)
	router.POST("/signup", h.Signup)
	router.GET("/login", h.Login)

	// Protected routes (Require authentication)
	protectedRoutes := router.Group("/")
	protectedRoutes.Use(m.Authenticate())

	// Admin Routes (Only Admins Can Access)
	adminRoutes := protectedRoutes.Group("/admin")
	adminRoutes.Use(m.RoleAuthMiddleware("admin")) // ✅ Only admin can access
	{
		adminRoutes.GET("/fetchAllUsers", h.FetchAllUsers)
		adminRoutes.PUT("/updateUserByID", h.UpdateUserByID)
		adminRoutes.DELETE("/deleteUserByID/:id", h.DeleteUserByID)
	}

	// User Routes (Users & Admins Can Access)
	userRoutes := protectedRoutes.Group("/user")
	userRoutes.Use(m.RoleAuthMiddleware("user", "admin")) // ✅ Both users & admins can access
	{
		userRoutes.GET("/fetchUserByID", h.FetchUserByID)
	}

	router.Run()
}

