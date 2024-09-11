package http

import (
	"encoding/json"
	"net/http"

	"github.com/datarohit/go-dynamodb-crud-api/utils/logger"
	"go.uber.org/zap"
)

type Response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result,omitempty"`
}

func NewResponse(data interface{}, status int) *Response {
	return &Response{
		Status: status,
		Result: data,
	}
}

func (resp *Response) Bytes() []byte {
	data, err := json.Marshal(resp)
	if err != nil {
		logger.GetLogger().Error("Failed to marshal response",
			zap.Error(err),
			zap.String("response_status", "500"),
			zap.String("response_result", "Internal Server Error"),
		)
		return []byte(`{"status":500, "result":"Internal Server Error"}`)
	}
	return data
}

func (resp *Response) String() string {
	return string(resp.Bytes())
}

func (resp *Response) SendResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(resp.Status)
	if _, err := w.Write(resp.Bytes()); err != nil {
		logger.GetLogger().Error("Failed to write response",
			zap.Error(err),
		)
	}
	logger.GetLogger().Info("Response sent",
		zap.Int("status", resp.Status),
		zap.Any("result", resp.Result),
	)
}

func StatusOK(w http.ResponseWriter, data interface{}) {
	NewResponse(data, http.StatusOK).SendResponse(w)
}

func StatusNoContent(w http.ResponseWriter) {
	NewResponse(nil, http.StatusNoContent).SendResponse(w)
}

func StatusBadRequest(w http.ResponseWriter, err error) {
	data := map[string]interface{}{"error": err.Error()}
	NewResponse(data, http.StatusBadRequest).SendResponse(w)
}

func StatusNotFound(w http.ResponseWriter, err error) {
	data := map[string]interface{}{"error": err.Error()}
	NewResponse(data, http.StatusNotFound).SendResponse(w)
}

func StatusMethodNotAllowed(w http.ResponseWriter) {
	NewResponse(nil, http.StatusMethodNotAllowed).SendResponse(w)
}

func StatusConflict(w http.ResponseWriter, err error) {
	data := map[string]interface{}{"error": err.Error()}
	NewResponse(data, http.StatusConflict).SendResponse(w)
}

func StatusInternalServerError(w http.ResponseWriter, err error) {
	data := map[string]interface{}{"error": err.Error()}
	NewResponse(data, http.StatusInternalServerError).SendResponse(w)
}
