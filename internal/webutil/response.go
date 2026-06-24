package webutil

import (
	"encoding/json"
	"net/http"
)

func Response(writer http.ResponseWriter, statusCode int, response interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, "Erro interno", http.StatusInternalServerError)
	}
}
