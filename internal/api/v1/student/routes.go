package student

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("students")
	h := NewHandler()
	r.POST("/", h.CreateStudent)
	r.GET("/", h.GetAllStudents)
	r.GET(":id", h.GetStudent)
	r.PUT(":id", h.UpdateStudent)
	r.DELETE(":id", h.DeleteStudent)
}
