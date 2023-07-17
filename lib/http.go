package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

func HTTPResponse(data interface{}, writer http.ResponseWriter, statusCode int) {
	respBody := make(map[string]interface{})
	respBody["result"] = data
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	err := json.NewEncoder(writer).Encode(&respBody)
	if err != nil {
		// Handle the error
		log.Printf(err.Error())
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
