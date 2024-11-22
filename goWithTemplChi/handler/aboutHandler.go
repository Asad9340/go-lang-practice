package handler

import (
	"goWithTemplChi/template"
	"log"
	"net/http"
)

func AboutUsHandler(w http.ResponseWriter, r *http.Request) {
	if	err:= template.AboutUs().Render(r.Context(),w); err !=nil{
		log.Fatal("Failed to load About us page")
	}

}