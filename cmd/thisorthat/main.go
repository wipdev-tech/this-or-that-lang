package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"wipdev-tech/this-or-that-lang/internal/database"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type service struct {
	db *database.Queries
}

func main() {
	dbURL, dbToken, err := getDBEnv()
	if err != nil {
		log.Fatal(err)
	}

	connURL := fmt.Sprintf("%s?authToken=%s", dbURL, dbToken)
	db, err := sql.Open("libsql", connURL)
	if err != nil {
		log.Fatal(err)
	}

	s := service{db: database.New(db)}

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
		"templates/top.html",
		"templates/bottom.html",
	))

	langs, err := s.db.GetLanguages(r.Context())
	if err != nil {
		log.Fatal(err)
	}
	rand.Shuffle(len(langs), func(i, j int) { langs[i], langs[j] = langs[j], langs[i] })

	comments := []string{
		"// but for languages!",
		"/* but for languages! */",
		"# but for languages!",
		"-- but for languages!",
		"; but for languages!",
		"(* but for languages! *)",
	}
	comment := comments[rand.Intn(len(comments))]

	tmpl.Execute(w, struct {
		Langs   []database.Language
		Comment string
	}{
		Langs:   langs,
		Comment: comment,
	})
}

func (s *service) handleRegister(w http.ResponseWriter, r *http.Request) {
}

func (s *service) handleLogin(w http.ResponseWriter, r *http.Request) {
}
