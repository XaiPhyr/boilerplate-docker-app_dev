package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	router "api_service_1/routers"
)

func main() {
	r := router.InitRouter()

	port := os.Getenv("HTTP_PORT")
	version := os.Getenv("VERSION")
	service := os.Getenv("SERVICE")

	log.SetFlags(log.Llongfile | log.LstdFlags)

	fmt.Println()
	log.Printf("-> Service: %s", service)
	log.Printf("-> Version: %s", version)
	log.Printf("-> Local:   http://localhost:%s", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	log.Printf("Error ListenAndServe: %s", err)
}
