package middlewares

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"

	"api_service_1/models"
	"api_service_1/services"
	s "api_service_1/services"
)

type Middleware struct {
	c services.Cache
}

var (
	LimiterExpiration = 1 * time.Minute
	RequestID         = middleware.RequestID
	RealIP            = middleware.RealIP
	Logger            = middleware.Logger
	Recoverer         = middleware.Recoverer
	GetHead           = middleware.GetHead
	HttpRate          = httprate.LimitByIP(5, LimiterExpiration)
)

func (m Middleware) Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Header.Get("User")

		t := m.c.GetCache(user)

		if err := s.VerifyJWT(t, w); err != nil {
			errObj := &models.ErrorObject{
				Code:    http.StatusUnauthorized,
				Message: "Token Expired",
			}

			jsonMarshal, _ := json.MarshalIndent(errObj, "", "  ")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(jsonMarshal))

			return
		}

		h.ServeHTTP(w, r)
	})
}

func (Middleware) UseMiddlewares(r *chi.Mux) {
	r.Use(RequestID)
	r.Use(RealIP)
	r.Use(Logger)
	r.Use(Recoverer)
	r.Use(GetHead)
	r.Use(HttpRate)
}
