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

	if request.Method != "POST" {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(`{"message": "Not found"}`))
	}

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
	if request.Method != "GET" {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(`{"message": "Not found"}`))
	}

	vars := mux.Vars(request)

	var url = database.FindUrlByShortcut(vars["shortcut"])

	// writer.WriteHeader(http.StatusMovedPermanently)
    // writer.Header().Add("Location", url)
	// writer.Header().Add("Content-Type", "text/html")

	// writer.Write([]byte(fmt.Sprintf(`
	// 	<html>
	// 	<head>
	// 		<title>Shorty</title>
	// 		</head>
	// 		<body>
	// 		<h1>Moved</h1>
	// 		<p>This page has moved to <a href="%s">%s</a>.</p>
	// 		</body>
	// 	</html>
	// `, url, url)))

	http.Redirect(writer, request, url, http.StatusMovedPermanently)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", CreateShortcutHandler)
	router.HandleFunc("/{shortcut}", AccessShortcutHandler)

	fmt.Println("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
