package controllers

import (
	"app_service_0/models"
	"app_service_0/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type AppController struct {
	c services.Cache
}

var endpoint = os.Getenv("ENDPOINT")

func (ac AppController) InitAppController(r *chi.Mux) {
	r.Route(fmt.Sprintf("%s/app-controller", endpoint), func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var data models.RedisPubSub

			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				log.Printf("Error JSON Decode: %s", err)
			}

			ac.c.RedisPub(data.Message)
			w.Write([]byte("INIT APP CONTROLLER!..."))
		})
	})
}
