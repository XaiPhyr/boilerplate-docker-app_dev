package middlewares

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
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

func (Middleware) UseMiddlewares(r *chi.Mux) {
	r.Use(RequestID)
	r.Use(RealIP)
	r.Use(Logger)
	r.Use(Recoverer)
	r.Use(GetHead)
	r.Use(HttpRate)
}
