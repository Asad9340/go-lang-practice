package main

import (
	"go-crud-mysql/config"
	"go-crud-mysql/router"
	"log"
	"net/http"
)

func main() {
	db := config.InitDatabase()
	r:= router.SetupRouter(db)
	log.Println("Server running at http://localhost:8080")
	if err:=http.ListenAndServe(":8080", r); err!=nil{
		log.Fatal("Failed to start server")
	}
}