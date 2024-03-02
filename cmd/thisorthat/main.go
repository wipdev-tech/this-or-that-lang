package main

import (
	"fmt"
	"net/http"
)

type service struct {
}

func main() {
	s := service{}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", s.handleHome)
	http.HandleFunc("/register", s.handleRegister)
	http.HandleFunc("/login", s.handleLogin)
	http.HandleFunc("/play", s.handlePlay)

	fmt.Println("This or That Lang server, let's go!")
}

func (s *service) handleHome(w http.ResponseWriter, r *http.Request) {
}

func (s *service) handleRegister(w http.ResponseWriter, r *http.Request) {
}

func (s *service) handleLogin(w http.ResponseWriter, r *http.Request) {
}

func (s *service) handlePlay(w http.ResponseWriter, r *http.Request) {
}
