package routes

import (
	"fmt"
	"log"
	"net/http"
)

type MediaHandler struct{}

func LogRequest(writer http.ResponseWriter, request *http.Request) {
	log.Print(request.URL.Path)
	_, err := fmt.Fprint(writer, request.URL.RawPath)
	fmt.Print("Came to Logger API")
	if err != nil {
		return
	}
}

func (h *MediaHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Print(request)
	if request.Method == "POST" {
		LogRequest(writer, request)
	} else {
		http.Error(writer, "method not allowed", http.StatusBadGateway)
		return
	}
}
