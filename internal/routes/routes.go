package routes

import (
	"net/http"
	"wipdev-tech/this-or-that-lang/internal/service"
)

func NewRouter(s service.Service) *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", s.HandleHome)
	mux.HandleFunc("/register", s.HandleRegister)
	mux.HandleFunc("/login", s.HandleLogin)
	mux.HandleFunc("/play", s.HandlePlay)

	return mux
}
