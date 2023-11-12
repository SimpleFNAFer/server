package home

import "net/http"

type Handler interface {
	Home(w http.ResponseWriter, r *http.Request)
}
