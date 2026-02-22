package server

import (
	"github.com/gin-gonic/gin"
	"github.com/green-house-grow-agro/gateway/internal/config"
)

func New(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	RegisterRoutes(router, cfg)
	return router
}

func RegisterRoutes(router *gin.Engine, cfg *config.Config) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.GET("/ready", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ready"})
	})
}
