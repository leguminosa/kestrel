package wrapper

import (
	"net/http"

	"github.com/leguminosa/kestrel/pkg/jsonx"
)

func OK(w http.ResponseWriter, data interface{}) {
	_ = wrapJSON(w, "success", http.StatusOK, data)
}

func BadRequest(w http.ResponseWriter, err error, data interface{}) {
	_ = wrapJSON(w, errMsg(err), http.StatusBadRequest, data)
}

func InternalServerError(w http.ResponseWriter, err error, data interface{}) {
	_ = wrapJSON(w, errMsg(err), http.StatusInternalServerError, data)
}

func errMsg(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func wrapJSON(w http.ResponseWriter, message string, statusCode int, data interface{}) error {
	return JSON(w, statusCode, constructResponse(message, statusCode, data))
}

func constructResponse(message, status, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"status":  status,
		"data":    data,
	}
}

func JSON(w http.ResponseWriter, statusCode int, msg interface{}) error {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	return jsonx.GetClient().Encode(w, msg)
}
