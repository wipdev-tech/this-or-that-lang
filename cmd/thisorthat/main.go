package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type service struct {
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	s := service{}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", s.handleHome)
	http.HandleFunc("/register", s.handleRegister)
	http.HandleFunc("/login", s.handleLogin)
	http.HandleFunc("/play", s.handlePlay)

	fmt.Println("This or That Lang server, let's go!")
	if os.Getenv("ENV") == "dev" {
		fmt.Println("Dev server started and running at http://localhost:8080")
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
	}
	fmt.Println("Server started and running")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+os.Getenv("PORT"), nil))
}

func (s *service) handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templates/index.html",
		// "templates/fragments.html",
	))

	tmpl.Execute(w, nil)
}

func (s *service) handleRegister(w http.ResponseWriter, r *http.Request) {
}

func (s *service) handleLogin(w http.ResponseWriter, r *http.Request) {
}

func (s *service) handlePlay(w http.ResponseWriter, r *http.Request) {
}
