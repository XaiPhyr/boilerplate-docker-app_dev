package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
	AppController
}

func (u User) InitUserController(r *chi.Mux) {
	r.Route(fmt.Sprintf("%s/users", endpoint), func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

			u.c.RedisPub("HELLO USERS")
			w.Write([]byte("Hello Users"))
		})
	})
}
