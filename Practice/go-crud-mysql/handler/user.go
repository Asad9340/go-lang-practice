// handler/user.go
package handler

import (
	"encoding/json"
	"go-crud-mysql/model"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)


func CreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var user model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if result := db.Create(&user); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}


func GetUsers(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var users []model.User
	if result := db.Find(&users); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	params := mux.Vars(r)
	id := params["id"]
	var user model.User
	if result := db.First(&user, id); result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}


func UpdateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	params := mux.Vars(r)
	id := params["id"]
	var user model.User
	if result := db.First(&user, id); result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if result := db.Save(&user); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter , r *http.Request, db *gorm.DB)  {
	params:=mux.Vars(r)
	id:=params["id"]
	var user model.User
	if result:=db.First(&user,id); result.Error !=nil{
		http.Error(w, "User not found", http.StatusNotFound)
    return
	}
	if result:=db.Delete(&user); result.Error !=nil{
		http.Error(w, result.Error.Error(),http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}