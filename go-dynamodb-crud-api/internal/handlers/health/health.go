package health

import (
	"errors"
	"net/http"
	"time"

	"github.com/datarohit/go-dynamodb-crud-api/internal/handlers"
	"github.com/datarohit/go-dynamodb-crud-api/internal/repository/adapter"
	httpUtils "github.com/datarohit/go-dynamodb-crud-api/utils/http"
	"github.com/datarohit/go-dynamodb-crud-api/utils/logger"
	"go.uber.org/zap"
)

type Handler struct {
	handlers.Interface
	Repository adapter.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Repository: repository,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if !h.Repository.Health() {
		msg := "Relational database not alive"
		logger.GetLogger().Error("Health check failed",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("message", msg),
			zap.Time("timestamp", time.Now()),
		)
		httpUtils.StatusInternalServerError(w, errors.New(msg))
		return
	}

	logger.GetLogger().Info("Health check successful",
		zap.String("method", r.Method),
		zap.String("url", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("message", "Service OK"),
		zap.Time("timestamp", time.Now()),
	)
	httpUtils.StatusOK(w, "Service OK")
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	logger.GetLogger().Warn("Unsupported method",
		zap.String("method", r.Method),
		zap.String("url", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("message", "Method Not Allowed"),
		zap.Time("timestamp", time.Now()),
	)
	httpUtils.StatusMethodNotAllowed(w)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	logger.GetLogger().Warn("Unsupported method",
		zap.String("method", r.Method),
		zap.String("url", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("message", "Method Not Allowed"),
		zap.Time("timestamp", time.Now()),
	)
	httpUtils.StatusMethodNotAllowed(w)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	logger.GetLogger().Warn("Unsupported method",
		zap.String("method", r.Method),
		zap.String("url", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("message", "Method Not Allowed"),
		zap.Time("timestamp", time.Now()),
	)
	httpUtils.StatusMethodNotAllowed(w)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	logger.GetLogger().Info("OPTIONS request received",
		zap.String("method", r.Method),
		zap.String("url", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.Time("timestamp", time.Now()),
	)
	httpUtils.StatusNoContent(w)
}
