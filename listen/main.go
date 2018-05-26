package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/dlockamy/videotogo"
)

var products = []Video{
	Video{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description: "Do cool things on your hover board and laser cannon"},
	Video{Id: 2, Name: "Laser Shooters", Slug: "laser-shooters", Description: "Do cool things with laser cannon"},
	Video{Id: 3, Name: "Hover Boat", Slug: "hover-boat", Description: "Do cool things on your boad and laser cannon"},
}

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Api is up and running"))
})

var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

var AddFeedBackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var product Video
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			product = p
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if product.Slug != "" {
		payload, _ := json.Marshal(product)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Product Not Found"))
	}
})

func main() {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./www/")))
	http.ListenAndServe(":3001", handlers.LoggingHandler(os.Stdout, r))
}
