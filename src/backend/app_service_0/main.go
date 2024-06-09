package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	router "app_service_0/routers"
	sql "app_service_0/sql"
	utils "app_service_0/utils"

	_ "github.com/lib/pq"
)

func main() {
	sql.MigrateSQL()
	r := router.InitRouter()

	version := os.Getenv("VERSION")
	port := os.Getenv("HTTP_PORT")
	service := os.Getenv("SERVICE")

	log.SetFlags(log.Llongfile | log.LstdFlags)

	fmt.Println()
	log.Printf("-> Service: %s", service)
	log.Printf("-> Version: %s", version)
	log.Printf("-> Local:   http://localhost:%s", port)

	utils.ParseFrontend(r)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	log.Printf("Error ListenAndServe: %s", err)
}
