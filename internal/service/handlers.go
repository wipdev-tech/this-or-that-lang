package service

import (
	"fmt"
	"html/template"
	"net/http"
	"wipdev-tech/this-or-that-lang/internal/database"
)

func (s *Service) HandleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/top.html",
		"templates/bottom.html",
	))

	tmpl.Execute(w, struct {
		Langs   []database.Language
		Comment string
	}{
		Langs:   s.getLangs(r.Context()),
		Comment: generateComment(),
	})
}

func (s *Service) HandleRegister(w http.ResponseWriter, r *http.Request) {
}

func (s *Service) HandleLogin(w http.ResponseWriter, r *http.Request) {
}

func (s *Service) HandlePlay(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templates/play.html",
		"templates/top.html",
		"templates/bottom.html",
	))

	votes, err := s.DB.GetAllVotes(r.Context())
	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, struct {
		Comment string
		Vote    database.GetAllVotesRow
	}{
		Comment: generateComment(),
		Vote:    votes[0],
	})
}
