package main

import (
	"log"
	"sre-bootcamp/internal/api/v1"
	"sre-bootcamp/internal/db"
	"sre-bootcamp/internal/service"
)

func main() {
	db.Init()
	studentService := service.NewStudentService()
	r := api.SetupRouter(studentService)
	log.Fatal(r.Run(":8080"))
}
