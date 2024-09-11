package logger

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestInitializeLogger(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Unexpected panic during logger initialization: %v", r)
		}
	}()
	InitializeLogger()
	if logger == nil {
		t.Error("Expected logger to be initialized, got nil")
	}
}

func TestLogRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test-url", nil)
	req.RemoteAddr = "127.0.0.1"

	core, logs := observer.New(zapcore.InfoLevel)
	testLogger := zap.New(core)

	logger = testLogger
	LogRequest(req, "test-trace-id")

	allLogs := logs.All()
	if len(allLogs) != 1 {
		t.Fatalf("Expected 1 log entry, got %d", len(allLogs))
	}

	entry := allLogs[0]
	if entry.Message != "incoming request" {
		t.Errorf("Expected log message 'incoming request', got %s", entry.Message)
	}

	fields := entry.ContextMap()
	if fields["method"] != "GET" {
		t.Errorf("Expected method 'GET', got %v", fields["method"])
	}
	if fields["url"] != "/test-url" {
		t.Errorf("Expected URL '/test-url', got %v", fields["url"])
	}
	if fields["remote_addr"] != "127.0.0.1" {
		t.Errorf("Expected remote address '127.0.0.1', got %v", fields["remote_addr"])
	}
	if fields["trace_id"] != "test-trace-id" {
		t.Errorf("Expected trace ID 'test-trace-id', got %v", fields["trace_id"])
	}
}

func TestLogError(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test-url", nil)
	req.RemoteAddr = "127.0.0.1"
	err := errors.New("test error")

	core, logs := observer.New(zapcore.ErrorLevel)
	testLogger := zap.New(core)

	logger = testLogger
	LogError(req, "test-trace-id", "test error message", err)

	allLogs := logs.All()
	if len(allLogs) != 1 {
		t.Fatalf("Expected 1 log entry, got %d", len(allLogs))
	}

	entry := allLogs[0]
	if entry.Message != "request error" {
		t.Errorf("Expected log message 'request error', got %s", entry.Message)
	}

	fields := entry.ContextMap()
	if fields["method"] != "GET" {
		t.Errorf("Expected method 'GET', got %v", fields["method"])
	}
	if fields["url"] != "/test-url" {
		t.Errorf("Expected URL '/test-url', got %v", fields["url"])
	}
	if fields["remote_addr"] != "127.0.0.1" {
		t.Errorf("Expected remote address '127.0.0.1', got %v", fields["remote_addr"])
	}
	if fields["trace_id"] != "test-trace-id" {
		t.Errorf("Expected trace ID 'test-trace-id', got %v", fields["trace_id"])
	}
	if fields["message"] != "test error message" {
		t.Errorf("Expected message 'test error message', got %v", fields["message"])
	}
	if fields["error"] != "test error" {
		t.Errorf("Expected error 'test error', got %v", fields["error"])
	}
}
