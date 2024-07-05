package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"

	"gateway/models"
	s "gateway/services"
)

type Middleware struct{}

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
		auth := r.Header.Get("Authorization")

		if auth == "" {
			errObj := &models.ErrorObject{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			}

			jsonMarshal, _ := json.MarshalIndent(errObj, "", "  ")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(jsonMarshal))

			return
		}

		t := strings.Split(auth, " ")[1]

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
