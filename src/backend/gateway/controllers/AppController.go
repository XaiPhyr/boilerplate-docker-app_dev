package controllers

import (
	"encoding/json"
	"fmt"
	"gateway/models"
	"gateway/services"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type AppController struct {
	c         services.Cache
	userModel *models.Users
}

var endpoint = os.Getenv("ENDPOINT")

var messageChan chan string

func (ac AppController) InitAppController(r *chi.Mux) {
	r.Route(fmt.Sprintf("%s/app-controller", endpoint), func(r chi.Router) {
		r.Get("/sse", ServerSent)

		r.Post("/pubsub", func(w http.ResponseWriter, r *http.Request) {
			var data services.RedisPubSub

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

func (a AppController) toJson(w http.ResponseWriter, b interface{}) {
	jsonMarshal, _ := json.MarshalIndent(b, "", "  ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonMarshal))
}

func (a AppController) handleError(w http.ResponseWriter, code int, message string) {
	errObj := models.ErrorObject{
		Code:    code,
		Message: message,
	}

	jsonMarshal, _ := json.MarshalIndent(errObj, "", "  ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(errObj.Code)
	w.Write([]byte(jsonMarshal))
}
