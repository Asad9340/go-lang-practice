package handler

import (
	"go-crud-mysql/test/template"
	"log"
	"net/http"
)

type RootHandler struct{}

func (rh *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func (rh *RootHandler) AboutData(w http.ResponseWriter, r *http.Request)   {
	if err:=template.AboutUs().Render(r.Context(),w); err!=nil{
		log.Fatal("Failed to load about us page")
	}
}