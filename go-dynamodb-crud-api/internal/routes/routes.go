package routes

import (
	"net"
	"net/http"
	"time"

	ServerConfig "github.com/datarohit/go-dynamodb-crud-api/config"
	HealthHandler "github.com/datarohit/go-dynamodb-crud-api/internal/handlers/health"
	ProductHandler "github.com/datarohit/go-dynamodb-crud-api/internal/handlers/product"
	"github.com/datarohit/go-dynamodb-crud-api/internal/repository/adapter"
	"github.com/datarohit/go-dynamodb-crud-api/utils/logger"
	"github.com/go-chi/chi/v5"
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
	r.EnableLogging()
	r.EnableTimeout()
	r.EnableRecovery()
	r.EnableRequestID()
	r.EnableRealIP()
}

func (r *Router) RouterHealth(repository adapter.Interface) {
	handler := HealthHandler.NewHandler(repository)

	r.router.Route("/health", func(route chi.Router) {
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Put("/", handler.Put)
		route.Delete("/", handler.Delete)
		route.Options("/", handler.Options)
	})
}

func (r *Router) RouterProduct(repository adapter.Interface) {
	handler := ProductHandler.NewHandler(repository)

	r.router.Route("/product", func(route chi.Router) {
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Get("/{ID}", handler.Get)
		route.Put("/{ID}", handler.Put)
		route.Delete("/{ID}", handler.Delete)
		route.Options("/", handler.Options)
	})
}

func (r *Router) EnableCORS() *Router {
	r.router.Use(r.config.Cors)
	return r
}

func (r *Router) EnableLogging() *Router {
	r.router.Use(loggingMiddleware)
	return r
}

func (r *Router) EnableTimeout() *Router {
	r.router.Use(timeoutMiddleware)
	return r
}

func (r *Router) EnableRecovery() *Router {
	r.router.Use(recoveryMiddleware)
	return r
}

func (r *Router) EnableRequestID() *Router {
	r.router.Use(requestIDMiddleware)
	return r
}

func (r *Router) EnableRealIP() *Router {
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
