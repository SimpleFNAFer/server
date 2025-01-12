package home

import "net/http"

type Handler interface {
	Home(w http.ResponseWriter, r *http.Request)
	Admin(w http.ResponseWriter, r *http.Request)
	AdminLoginGet(w http.ResponseWriter, _ *http.Request)
	AdminLoginPost(w http.ResponseWriter, r *http.Request)
	Block(w http.ResponseWriter, r *http.Request)
	Unblock(w http.ResponseWriter, r *http.Request)
	Forbidden(w http.ResponseWriter, _ *http.Request)
}
