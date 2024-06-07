package router

import (
	"app_service_0/controllers"
	"app_service_0/middlewares"

	"github.com/go-chi/chi/v5"
)

type Router struct{}

var (
	m             middlewares.Middleware
	appController = &controllers.AppController{}
)

func InitRouter() (r *chi.Mux) {
	r = chi.NewRouter()

	m.UseMiddlewares(r)

	// @router /app-controller
	appController.InitAppController(r)

	return
}
