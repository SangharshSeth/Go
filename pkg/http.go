package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

func SuccessStatusCode(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}

func HTTPResponse(data interface{}, writer http.ResponseWriter, statusCode int, headers map[string]string) {
	log.Print(headers)
	log.Print("In Headers")
	respBody := make(map[string]interface{})
	if SuccessStatusCode(statusCode) {
		respBody["status"] = "Success"
	} else {
		respBody["status"] = "Failed"
	}
	respBody["result"] = data
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
