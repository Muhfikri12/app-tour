package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Middleware struct {
	Log *zap.Logger
}

func NewMiddleware(log *zap.Logger) Middleware {
	return Middleware{
		Log: log,
	}
}

func (middleware *Middleware) MiddlewareLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		handler.ServeHTTP(w, r)

		duration := time.Since(start)

		middleware.Log.Info("http request", zap.String("url", r.URL.String()), zap.Duration("duration", duration))
	})

}