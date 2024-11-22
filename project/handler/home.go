package handler

import (
	"net/http"

	"goWithTemplChi/templates"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	users := []templates.User{
		{Name: "John Doe", Email: "john@example.com"},
		{Name: "Jane Doe", Email: "jane@example.com"},
	}
	templates.Layout().Render(r.Context(), w, templates.Home(users))
}
