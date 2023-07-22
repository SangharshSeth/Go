package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

func HTTPResponse(data interface{}, writer http.ResponseWriter, statusCode int, headers map[string]string) {
	log.Print(headers)
	respBody := map[string]interface{}{
		"result": data,
		"status": http.StatusText(statusCode),
	}

	for key, value := range headers {
		writer.Header().Set(key, value)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	err := json.NewEncoder(writer).Encode(&respBody)
	if err != nil {
		log.Print(err.Error())
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
