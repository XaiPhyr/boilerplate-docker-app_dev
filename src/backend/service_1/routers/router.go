package router

import (
	"api_service_1/controllers"
	"api_service_1/middlewares"
	"api_service_1/models"
	"encoding/json"
	"log"
	"net/http"
	"os"

	utils "api_service_1/utils"

	"github.com/go-chi/chi/v5"
)

type Router struct{}

var (
	m                 middlewares.Middleware
	productController = &controllers.Products{}
)

func InitRouter() (r *chi.Mux) {
	endpoint := os.Getenv("ENDPOINT")

	r = chi.NewRouter()

	m.UseMiddlewares(r)

	var mux = models.MuxServer{
		Mux:      r,
		Endpoint: endpoint,
	}

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
