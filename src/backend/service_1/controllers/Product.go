package controllers

import (
	"api_service_1/models"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Products struct {
	AppController
}

func (p Products) InitProductController(m models.MuxServer) {
	m.Mux.Route(fmt.Sprintf("%s/products", m.Endpoint), func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(p.mw.Authenticate)

			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Product from API Service 01"))
			})
		})
	})
}
