package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type HttpHandler struct {
}

func getInfo(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println("Failed to restrict File Size")
	}
	file, fileHeader, err := request.FormFile("Document")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			panic(err.Error())
		}
	}(file)

	//store the file locally
	_, err = os.Stat("./uploads")
	if err != nil {
		if !os.IsExist(err) {
			fileErr := os.Mkdir("./uploads", os.ModePerm)
			if fileErr != nil {
				http.Error(writer, fileErr.Error(), http.StatusInternalServerError)
			}
		}
	}
	randomName := rand.Int()
	fmt.Println(randomName)
	dst, _ := os.Create(fmt.Sprintf("./uploads/%d%s", randomName, filepath.Ext(fileHeader.Filename)))
	curFile := strconv.Itoa(randomName) + filepath.Ext(fileHeader.Filename)
	fmt.Println("Current File Created in New Directory", curFile)
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			panic(err.Error())
		}
	}(dst)

	newFilePath := "./uploads/" + string(curFile)
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	//get bytes from the file
	fileBytes, _ := os.ReadFile(newFilePath)

	fmt.Println(fileBytes)
	writer.WriteHeader(http.StatusAccepted)
	writer.Header().Set("Content-Type", "application/json")

	fileDetails := make(map[string]string)
	fileDetails["fileName"] = fileHeader.Filename
	fileDetails["fileSize"] = strconv.FormatInt(fileHeader.Size, 10)
	fileDetails["uploadStatus"] = "FAILED"

	respErr := json.NewEncoder(writer).Encode(&fileDetails)
	if err != nil {
		panic(respErr.Error())
	}
}

func (h *HttpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request)
	if request.Method == "POST" {
		writer.Header().Set("Content-Type", "application/json")
		getInfo(writer, request)
	} else {
		http.Error(writer, "method not allowed", http.StatusBadGateway)
		return
	}
}
