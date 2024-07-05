package controllers

import (
	"fmt"
	"gateway/models"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
)

type Products struct {
	AppController
}

func (p Products) InitProductController(m models.MuxServer) {
	m.Mux.Route(fmt.Sprintf("%s/products", m.Endpoint), func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				origin, err := url.Parse(m.Proxy)
				if err != nil {
					panic(err)
				}

				proxy := httputil.NewSingleHostReverseProxy(origin)
				proxy.ServeHTTP(w, r)
			})
		})
	})
}
