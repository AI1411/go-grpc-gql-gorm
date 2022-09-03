package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AI1411/go-grpc-gql/internal/db"
	"github.com/AI1411/go-grpc-gql/internal/infra/repository"
)

func main() {
	client, err := db.NewClient()
	if err != nil {
		panic(err)
	}

	repo := repository.NewArticleRepository(client)
	a, err := repo.ListArticles(context.Background())
	if err != nil {
		log.Printf("repo error: %v", err)
	}
	fmt.Printf("articles = %+v", a)
}
