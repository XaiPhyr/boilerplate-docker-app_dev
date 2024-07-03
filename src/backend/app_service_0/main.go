package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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

	r.HandleFunc("/reroutes", func(w http.ResponseWriter, r *http.Request) {
		remote, _ := url.Parse("http://app_service_1:8201/routed")
		proxy := httputil.NewSingleHostReverseProxy(remote)

		proxy.ServeHTTP(w, r)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	log.Printf("Error ListenAndServe: %s", err)
}
