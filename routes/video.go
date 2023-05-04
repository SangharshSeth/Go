package routes

import (
	"fmt"
	"log"
	"net/http"
)

type MediaHandler struct{}

func StreamMedia(writer http.ResponseWriter, request *http.Request) {

}

func LogRequest(writer http.ResponseWriter, request *http.Request) {
	log.Print(request.URL.Path)
	fmt.Fprint(writer, request.URL.RawPath)
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
