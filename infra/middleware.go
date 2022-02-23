package infra

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// DatabaseMiddleware for gin to pass DB context around
func DatabaseMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", db)
		c.Next()
	}
}

// RedisMiddleware for gin to pass DB context around
func RedisMiddleware(db *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("redis", db)
		c.Next()
	}
}
