package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		Warn(logrus.Fields{
			"action":     "WriteJSONResponse",
			"statusCode": statusCode,
		}, "Failed to write JSON response: "+err.Error())
	}
}

func ParseBody(r *http.Request, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		Warn(logrus.Fields{
			"action": "ParseBody",
			"url":    r.URL.Path,
		}, "Error reading request body: "+err.Error())
		return err
	}

	err = json.Unmarshal(body, x)
	if err != nil {
		Warn(logrus.Fields{
			"action": "ParseBody",
			"url":    r.URL.Path,
		}, "Error unmarshalling JSON: "+err.Error())
		return err
	}

	return nil
}
