package main

import (
	"log"
	"sre-bootcamp/internal/api/v1"
	"sre-bootcamp/internal/db"
)

func main() {
	db.Init()
	r := api.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
