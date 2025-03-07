package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ctxKey int

const Key ctxKey = 1

func (m *Middleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")

		// Splitting the Authorization header based on the space character.
		// Boats "Bearer" and the actual token
		parts := strings.Split(authHeader, " ")
		// Checking the format of the Authorization header
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			// If the header format doesn't match required format, log and send an error
			err := errors.New("expected authorization header format: Bearer <token>")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		claims, err := m.auth.ValidateToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired or invalid"})
			return
		}

		// Add role and email to context
		c.Set("role", claims.Role)
		c.Set("email", claims.Email)

		// Continue to the next middleware
		c.Next()

	}
}

func (m *Middleware) RoleAuthMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the role from the Gin context
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access denied: no role found"})
			return
		}

		// Convert to string (Type Assertion)
		roleStr, ok := role.(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid role type"})
			return
		}

		// Check if the user role is in the allowed roles
		for _, allowedRole := range requiredRoles {
			if roleStr == allowedRole {
				c.Next() // User is authorized, proceed to the next handler
				return
			}
		}

		// If no match, deny access
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access denied"})
	}
}
