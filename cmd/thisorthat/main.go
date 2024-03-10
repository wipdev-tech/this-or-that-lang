package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"wipdev-tech/this-or-that-lang/internal/database"
	"wipdev-tech/this-or-that-lang/internal/routes"
	"wipdev-tech/this-or-that-lang/internal/service"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

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

	s := service.New()
	s.DB = database.New(db)
	mux := routes.NewRouter(s)
	server := http.Server{Handler: mux}

	fmt.Println("This or That Lang server, let's go!")
	if os.Getenv("ENV") == "dev" {
		fmt.Println("Dev server started and running at http://localhost:8080")
		server.Addr = "localhost:8080"
	} else {
		fmt.Println("Server started and running")
		server.Addr = "0.0.0.0:" + os.Getenv("PORT")
	}
	log.Fatal(server.ListenAndServe())
}
