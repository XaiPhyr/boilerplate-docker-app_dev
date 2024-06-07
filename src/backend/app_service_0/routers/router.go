package router

import (
	"app_service_0/controllers"

	"github.com/go-chi/chi/v5"
)

type Router struct{}

var (
	appController = &controllers.AppController{}
)

func InitRouter() (r *chi.Mux) {
	r = chi.NewRouter()

	// @router /app-controller
	appController.InitAppController(r)

	return
}
