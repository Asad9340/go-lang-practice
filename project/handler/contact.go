package handler

import (
	"net/http"

	"goWithTemplChi/templates"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	templates.Layout().Render(r.Context(), w, templates.Contact())
}
