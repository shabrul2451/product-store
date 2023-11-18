package main

import (
	"Product-Store/api"
	"log"
)

func main() {
	r := api.SetupRoutes()
	err := r.Run(":8085")
	if err != nil {
		log.Fatal(err.Error())
	}
}
