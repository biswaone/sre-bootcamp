package api

import (
	"sre-bootcamp/internal/api/health"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/healthcheck", health.HealthCheck)

	return r
}
