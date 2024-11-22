package main

import (
	"log"
	"net/http"
	"project/config"
	"project/handler"
)

func main() {
	config.InitDB("postgres://user:password@localhost:5432/dbname")
	defer config.DB.Close()

	r := handler.Router()

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
