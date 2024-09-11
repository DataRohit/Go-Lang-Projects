package routes

import (
	"net"
	"net/http"
	"time"

	ServerConfig "github.com/datarohit/go-dynamodb-crud-api/config"
	"github.com/datarohit/go-dynamodb-crud-api/internal/repository/adapter"
	"github.com/datarohit/go-dynamodb-crud-api/utils/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig(30, 5).SetTimeout(ServerConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters(repository adapter.Interface) *chi.Mux {
	r.setConfigsRouters()
	return r.router
}

func (r *Router) setConfigsRouters() {
	r.EnableCORS()
	r.EnableLogger()
	r.EnableTimeout()
	r.EnableRecover()
	r.EnableRequestID()
	r.EnableRealIP()
	r.EnableCustomLogging()
	r.EnableCustomTimeout()
	r.EnableCustomRecovery()
	r.EnableCustomRequestID()
	r.EnableCustomRealIP()
}

func (r *Router) EnableLogger() *Router {
	r.router.Use(middleware.Logger)
	return r
}

func (r *Router) EnableTimeout() *Router {
	r.router.Use(middleware.Timeout(r.config.GetTimeout()))
	return r
}

func (r *Router) EnableCORS() *Router {
	r.router.Use(r.config.Cors)
	return r
}

func (r *Router) EnableRecover() *Router {
	r.router.Use(middleware.Recoverer)
	return r
}

func (r *Router) EnableRequestID() *Router {
	r.router.Use(middleware.RequestID)
	return r
}

func (r *Router) EnableRealIP() *Router {
	r.router.Use(middleware.RealIP)
	return r
}

func (r *Router) EnableCustomLogging() *Router {
	r.router.Use(loggingMiddleware)
	return r
}

func (r *Router) EnableCustomTimeout() *Router {
	r.router.Use(timeoutMiddleware)
	return r
}

func (r *Router) EnableCustomRecovery() *Router {
	r.router.Use(recoveryMiddleware)
	return r
}

func (r *Router) EnableCustomRequestID() *Router {
	r.router.Use(requestIDMiddleware)
	return r
}

func (r *Router) EnableCustomRealIP() *Router {
	r.router.Use(realIPMiddleware)
	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := uuid.New().String()
		logger.LogRequest(r, traceID)
		next.ServeHTTP(w, r)
	})
}

func timeoutMiddleware(next http.Handler) http.Handler {
	return http.TimeoutHandler(next, 30*time.Second, "Request timed out")
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := uuid.New().String()
		defer func() {
			if rec := recover(); rec != nil {
				errMsg := "Internal Server Error"
				logger.GetLogger().Error("request recovery",
					zap.String("method", r.Method),
					zap.String("url", r.URL.Path),
					zap.String("remote_addr", r.RemoteAddr),
					zap.String("trace_id", traceID),
					zap.String("message", errMsg),
					zap.Any("recovered", rec),
					zap.Time("timestamp", time.Now()),
				)
				http.Error(w, errMsg, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := uuid.New().String()
		r.Header.Set("X-Request-ID", traceID)
		next.ServeHTTP(w, r)
	})
}

func realIPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		realIP := r.Header.Get("X-Forwarded-For")
		if realIP == "" {
			realIP, _, _ = net.SplitHostPort(r.RemoteAddr)
		}
		r.Header.Set("X-Real-IP", realIP)
		next.ServeHTTP(w, r)
	})
}
