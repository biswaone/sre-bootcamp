package student

import (
	"sre-bootcamp/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, studentService service.IStudentService) {
	r := rg.Group("students")
	h := NewHandler(studentService)
	r.POST("/", h.CreateStudent)
	r.GET("/", h.GetAllStudents)
	r.GET("/:id", h.GetStudent)
	r.PUT("/:id", h.UpdateStudent)
	r.DELETE("/:id", h.DeleteStudent)
}
