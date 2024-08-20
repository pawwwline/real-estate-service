package utils

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Code      int    `json:"code"`
}

func ErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	requestID, ok := r.Context().Value(middleware.RequestIDKey).(string)
	if !ok {
		requestID = "unknown"
	}

	log.Printf("Request handling error: %v, request_id: %s", err, requestID)

	errorResponse := ErrorResponse{
		Message:   "что-то пошло не так",
		RequestID: requestID,
		Code:      http.StatusInternalServerError,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	if encodeErr := json.NewEncoder(w).Encode(errorResponse); encodeErr != nil {
		log.Printf("Failed to encode error response: %v", encodeErr)
	}
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	ErrorHandlerFunc(w, r, errors.New("Невалидные данные ввода\n"))
}

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	ErrorHandlerFunc(w, r, errors.New("Неавторизованный доступ"))
}

func InternalServerError(w http.ResponseWriter, r *http.Request, message string) {
	ErrorHandlerFunc(w, r, errors.New(message))
}
