package service

import (
	"context"
	"log"
	"math/rand"
	"wipdev-tech/this-or-that-lang/internal/database"
)

type Service struct {
	DB *database.Queries
}

func New() Service {
	return Service{}
}

func (s *Service) getLangs(ctx context.Context) []database.Language {
	langs, err := s.DB.GetLanguages(ctx)
	if err != nil {
		log.Fatal(err)
	}
	rand.Shuffle(len(langs), func(i, j int) { langs[i], langs[j] = langs[j], langs[i] })
	return langs
}

func generateComment() string {
	comments := []string{
		"// but for languages!",
		"/* but for languages! */",
		"# but for languages!",
		"-- but for languages!",
		"; but for languages!",
		"(* but for languages! *)",
	}
	comment := comments[rand.Intn(len(comments))]
	return comment
}
