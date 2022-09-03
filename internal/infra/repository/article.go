package repository

import (
	"context"
	pb "github.com/AI1411/go-grpc-gql/grpc"
	"github.com/AI1411/go-grpc-gql/internal/db"
)

type ArticleRepository struct {
	dbClient *db.Client
}

func (r *ArticleRepository) InsertArticle(ctx context.Context, input *pb.ArticleInput) (int64, error) {
	return 0, nil
}
