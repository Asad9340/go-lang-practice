package handler

import (
	"encoding/json"
	"go-crud-mysql/model"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) CreateUser (w http.ResponseWriter, r *http.Request ){
	var user model.User
	if err:=json.NewDecoder(r.Body).Decode(&user); err!=nil{
		http.Error(w, "Invalid Input", http.StatusBadRequest)
    return
	}

	if result := h.DB.Create(&user);result.Error != nil{
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
    return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request){
	var users []model.User
	if result:= h.DB.Find(&users); result.Error != nil{
		http.Error(w, result.Error.Error(),http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}


func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var user model.User
	if result := h.DB.First(&user, id); result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if result := h.DB.Save(&user); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}


func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var user model.User
	if result := h.DB.First(&user, id); result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if result := h.DB.Delete(&user); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
