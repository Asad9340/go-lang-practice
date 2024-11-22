package handler

import (
	"goWithTemplChi/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err:= template.Index().Render(r.Context(),w); err !=nil{
		log.Fatal("Failed to load index file: ")
	}
}