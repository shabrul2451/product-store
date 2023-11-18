package main

import (
	"Product-Store/api"
	"Product-Store/config"
	"log"
)

func main() {
	config.InitEnv()
	r := api.SetupRoutes()
	err := r.Run(":8085")
	if err != nil {
		log.Fatal(err.Error())
	}
}
