package handler

import (
	"net/http"

	"goWithTemplChi/templates"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	templates.Layout().Render(r.Context(), w, templates.About())
}
