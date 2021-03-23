package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"ha.si/shorty/database"
	"ha.si/shorty/generator"

	"github.com/gorilla/mux"
)

type CreateShortcutRequestBody struct {
	Url 	string `json:"url" binding:"required"`
	Format 	string `json:"format,omitempty"`
}

func CreateShortcutHandler(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")

	var body CreateShortcutRequestBody	

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&body)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"message": "Internal Server Error"}`))
		return
	}

	var shortcut = generator.GenerateMnemonic()

	database.AddShortcut(shortcut, body.Url)

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(shortcut))
}

func AccessShortcutHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	var url = database.FindUrlByShortcut(vars["shortcut"])

	if len(url) == 0 {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(writer, request, url, http.StatusMovedPermanently)
}

func main() {
	router := mux.NewRouter()

	createRouter := router.Methods("POST").Subrouter()
	createRouter.HandleFunc("/", CreateShortcutHandler)

	accessRouter := router.Methods("GET").Subrouter()
	accessRouter.HandleFunc("/{shortcut}", AccessShortcutHandler)

	fmt.Println("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
