package controllers

import (
	"api_service_1/middlewares"
	"api_service_1/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type AppController struct {
	mw middlewares.Middleware
}

var endpoint = os.Getenv("ENDPOINT")

func (a AppController) InitAppController(r *chi.Mux) {
	r.Route(fmt.Sprintf("%s/app-controller", endpoint), func(r chi.Router) {

	})
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
