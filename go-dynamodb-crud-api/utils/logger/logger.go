package logger

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	InitializeLogger()
	defer logger.Sync()
}

func InitializeLogger() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
}

func GetLogger() *zap.Logger {
	return logger
}

type LogData struct {
	Method     string      `json:"method"`
	URL        string      `json:"url"`
	RemoteAddr string      `json:"remote_addr"`
	Message    string      `json:"message,omitempty"`
	Error      string      `json:"error,omitempty"`
	StatusCode int         `json:"status_code,omitempty"`
	Timestamp  time.Time   `json:"timestamp"`
	TraceID    string      `json:"trace_id,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func LogRequest(r *http.Request, traceID string) {
	logger.Info("incoming request",
		zap.String("method", r.Method),
		zap.String("url", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("trace_id", traceID),
		zap.Time("timestamp", time.Now()),
	)
}

func LogError(r *http.Request, traceID, msg string, err error) {
	logger.Error("request error",
		zap.String("method", r.Method),
		zap.String("url", r.URL.Path),
		zap.String("remote_addr", r.RemoteAddr),
		zap.String("trace_id", traceID),
		zap.String("message", msg),
		zap.Error(err),
		zap.Time("timestamp", time.Now()),
	)
}
