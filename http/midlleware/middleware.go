package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/time/rate"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			status := http.StatusInternalServerError

			switch c.Errors.Last().Type {
			case gin.ErrorTypeBind:
				status = http.StatusBadRequest
			case gin.ErrorTypePublic:
				status = http.StatusBadRequest
			case gin.ErrorTypePrivate:
				status = http.StatusInternalServerError
			}

			c.JSON(status, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}
	}
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("JWT")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading cookie"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			secretKey := os.Getenv("JWT_SECRET")
			return []byte(secretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error parse token"})
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User unuthorized"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}
		role := claims["role"].(string)
		c.Set("role", role)
		c.Next()

		// c.Set("user_id", claims["user_id"])
		// c.Set("role", claims["role"])

	}
}

func requireRole(requileRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != requileRole {
			c.JSON(403, gin.H{"error": "Insufficient permissions"})
			c.Abort()
		}
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := map[string]bool{
		"http://localhost:3000": true,
		"https://myapp.com":     true,
	}

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		if allowedOrigins[origin] {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization, X-Requested-With, Accept",
		)
		c.Writer.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, PATCH, DELETE, OPTIONS",
		)
		c.Writer.Header().Set("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			log.Printf("Preflight request from %s", origin)
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime)
		clientIp := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()

		log.Printf("Request: %s %s from %s | Status: %d | Duration: %s", method, path, clientIp, statusCode, duration)
	}
}

func RateLimiter(l *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !l.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			return
		}
		c.Next()
	}
}
