package main

import (
	"fmt"
	"go-crud-mysql/handler"
	"go-crud-mysql/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// Initialize the database connection
func InitDatabase() {
	dsn := "root:9340@tcp(127.0.0.1:3308)/cruddb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Database connected successfully")
	db.AutoMigrate(&model.User{})
}

func main() {
	InitDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateUser(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handler.GetUsers(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.GetUser(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdateUser(w, r, db)
	}).Methods("PUT")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.DeleteUser(w, r, db)
	}).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
