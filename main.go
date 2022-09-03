package main

import (
	"fmt"
	"github.com/AI1411/go-grpc-gql/internal/db"
)

func main() {
	client, err := db.NewClient()
	if err != nil {
		panic(err)
	}
	fmt.Printf("client = %+v", client)
}
