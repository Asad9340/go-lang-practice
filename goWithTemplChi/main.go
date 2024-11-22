package main

import (
	"fmt"
	"goWithTemplChi/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
		r.Get("/", handler.HomeHandler)
		r.Get("/about", handler.AboutUsHandler)
	fmt.Println("Server is running on port: 5050")
	log.Fatal(http.ListenAndServe(":5050",r))
}