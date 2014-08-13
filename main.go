package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("/Users/pedro/Documents/Avalanche/Videos")))
	http.Handle("/", router)
	http.ListenAndServe(getPort(), nil)
}
