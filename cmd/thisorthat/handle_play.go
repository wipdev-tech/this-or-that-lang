package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func (s *service) handlePlay(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templates/play.html",
		"templates/top.html",
		"templates/bottom.html",
	))

	votes, err := s.db.GetAllVotes(r.Context())
	if err != nil {
		fmt.Println(err)
	}

	tmpl.Execute(w, votes[0])
}
