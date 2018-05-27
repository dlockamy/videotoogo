package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./www")))

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}
