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

var messageChan chan string

func (ac AppController) InitAppController(r *chi.Mux) {
	r.Route(fmt.Sprintf("%s/app-controller", endpoint), func(r chi.Router) {
		r.Get("/sse", ServerSent)

		r.Post("/pubsub", func(w http.ResponseWriter, r *http.Request) {
			var data models.RedisPubSub

			err := json.NewDecoder(r.Body).Decode(&data)
			if err != nil {
				log.Printf("Error JSON Decode: %s", err)
			}

			ac.c.RedisPub(data.Message)
			w.Write([]byte("INIT APP CONTROLLER!..."))

			messageChan <- data.Message
		})
	})
}

func ServerSent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher := w.(http.Flusher)
	messageChan = make(chan string)

	defer func() {
		close(messageChan)
		messageChan = nil
	}()

	for {
		select {
		case message := <-messageChan:
			fmt.Fprintf(w, "data: %s\n\n", message)
			flusher.Flush()

		case <-r.Context().Done():
			break
		}
	}
}
