package router

import (
	"encoding/json"
	"gateway/controllers"
	"gateway/middlewares"
	"gateway/models"
	"log"
	"net/http"
	"os"

	utils "gateway/utils"

	"github.com/go-chi/chi/v5"
)

type Router struct{}

var (
	m                        middlewares.Middleware
	appController            = &controllers.AppController{}
	authenticationController = &controllers.Authentication{}
	userController           = &controllers.Users{}
	productController        = &controllers.Products{}
)

func InitRouter() (r *chi.Mux) {
	proxy := os.Getenv("HTTP_PROXY")
	endpoint := os.Getenv("ENDPOINT")

	r = chi.NewRouter()

	m.UseMiddlewares(r)

	var mux = models.MuxServer{
		Mux:      r,
		Proxy:    proxy,
		Endpoint: endpoint,
	}

	// @router /app-controller
	appController.InitAppController(r)

	// @router /auth
	authenticationController.InitAuthenticationController(mux)

	// @router /users
	userController.InitUserController(mux)

	// @router /products
	productController.InitProductController(mux)

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
