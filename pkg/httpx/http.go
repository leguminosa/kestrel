package httpx

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	StatusCode int
	Data       interface{}
}

type HandlerFunc func(*http.Request) (*Response, error)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp, err := fn(r)
	if err != nil {
		if _, errWrite := w.Write(writeErrorResponse(w, err)); errWrite != nil {
			log.Printf("write http response error {%s}", errWrite)
		}

		log.Printf("error returned from the handler {%s}", err)

		return
	}

	w.WriteHeader(resp.StatusCode)

	httpResponse := make(map[string]interface{})
	httpResponse["data"] = resp.Data

	response, err := json.Marshal(httpResponse)
	if err != nil {
		if _, errWrite := w.Write(writeErrorResponse(w, err)); errWrite != nil {
			log.Printf("write http response error {%s}", errWrite)
		}

		return
	}

	_, err = w.Write(response)
	if err != nil {
		log.Printf("write http response error {%s}", err)
	}
}

func writeErrorResponse(w http.ResponseWriter, err error) []byte {
	w.Header().Set("Content-Type", "application/json")

	switch {
	case err != nil:
		w.WriteHeader(http.StatusInternalServerError)
		return buildErrorResponse(err.Error())
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return buildErrorResponse(http.StatusText(http.StatusInternalServerError))
	}
}

func buildErrorResponse(message string) []byte {
	errorResponse := make(map[string]interface{})
	errorResponse["message"] = message

	response, err := json.Marshal(errorResponse)
	if err != nil {
		return []byte(http.StatusText(http.StatusInternalServerError))
	}

	return response
}
