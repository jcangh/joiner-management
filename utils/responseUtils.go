package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseOk(response http.ResponseWriter, v any) {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(v)
}

func ResponseCreated(response http.ResponseWriter, v any) {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(v)
}

func ResponseNotFound(response http.ResponseWriter) {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusNotFound)
}

func ResponseNotContent(response http.ResponseWriter) {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusNoContent)
}

func ResponseError(response http.ResponseWriter, v any) {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(response).Encode(v)
}

func ResponseBadRequest(response http.ResponseWriter) {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusBadRequest)
}
