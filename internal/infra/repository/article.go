package repository

import (
	"context"

	pb "github.com/AI1411/go-grpc-gql/grpc"
	"github.com/AI1411/go-grpc-gql/internal/db"
)

type Article struct {
	ID      int64
	Title   string
	Author  string
	Content string
}

type ArticleRepository struct {
	dbClient *db.Client
}

func NewArticleRepository(dbClient *db.Client) *ArticleRepository {
	return &ArticleRepository{
		dbClient: dbClient,
	}
}

func (r *ArticleRepository) InsertArticle(ctx context.Context, input *pb.ArticleInput) (int64, error) {
	return 0, nil
}

func (r *ArticleRepository) SelectArticleByID(ctx context.Context, id int64) (*pb.Article, error) {
	// DBからIDに基づいて記事をSELECT
	return nil, nil
}

func (r *ArticleRepository) UpdateArticle(ctx context.Context, id int64, input *pb.ArticleInput) error {
	//　DB内の記事をUPDATE
	return nil
}

func (r *ArticleRepository) DeleteArticle(ctx context.Context, id int64) error {
	// DB内の記事をDELETE
	return nil
}

func (r *ArticleRepository) ListArticles(ctx context.Context) ([]Article, error) {
	// articlesテーブルの記事を全取得
	var article []Article
	if err := r.dbClient.Conn(ctx).Find(&article).Error; err != nil {
		return nil, err
	}
	return article, nil
}
