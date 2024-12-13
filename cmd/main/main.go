package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sayandas-sd/stocksApiServer/router"
)

func main() {
	r := router.Router()

	fmt.Println("server is running on port: 8080..")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server: %v", err)
	}

}
