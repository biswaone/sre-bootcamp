package api

import (
	"sre-bootcamp/internal/api/health"
	"sre-bootcamp/internal/api/v1/student"
	"sre-bootcamp/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(studentService service.IStudentService) *gin.Engine {
	r := gin.Default()

	r.GET("/healthcheck", health.HealthCheck)

	v1 := r.Group("/api/v1")
	{
		student.RegisterRoutes(v1, studentService)
	}

	return r
}
