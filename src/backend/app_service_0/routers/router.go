package router

import (
	"app_service_0/controllers"
	"app_service_0/middlewares"
	"encoding/json"
	"log"
	"net/http"

	utils "app_service_0/utils"

	"github.com/go-chi/chi/v5"
)

type Router struct{}

var (
	m              middlewares.Middleware
	appController  = &controllers.AppController{}
	userController = &controllers.User{}
)

func InitRouter() (r *chi.Mux) {
	r = chi.NewRouter()

	m.UseMiddlewares(r)

	// @router /app-controller
	appController.InitAppController(r)

	// @router /users
	userController.InitUserController(r)

	// @status 404, 405
	PageNotFound(r)
	MethodNotAllowed(r)

	return
}

func PageNotFound(r *chi.Mux) {
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		content, err := utils.ParseHTML("templates/404.html", nil)

		if err != nil {
			log.Printf("Error: %s", err)
			return
		}

		w.Write([]byte(string(content)))
	})
}

func MethodNotAllowed(r *chi.Mux) {
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		b := map[string]interface{}{
			"err":  "method not allowed",
			"code": http.StatusMethodNotAllowed,
		}

		jsonMarshal, _ := json.MarshalIndent(b, "", "  ")

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(jsonMarshal))
	})
}
