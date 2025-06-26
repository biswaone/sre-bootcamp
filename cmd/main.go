package main

import (
	"log"
	"sre-bootcamp/internal/api/v1"
)

func main() {
	r := api.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
