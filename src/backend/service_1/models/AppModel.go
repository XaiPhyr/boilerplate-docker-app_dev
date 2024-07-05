package models

import (
	"github.com/go-chi/chi/v5"
)

type (
	MuxServer struct {
		Mux      *chi.Mux
		Endpoint string
		Proxy    string
	}

	ErrorObject struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)
