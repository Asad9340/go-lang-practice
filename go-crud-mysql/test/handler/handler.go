package handler

import "net/http"

type HandlerInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	AboutData(w http.ResponseWriter , r *http.Request)
}