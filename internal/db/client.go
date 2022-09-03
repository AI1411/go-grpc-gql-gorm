package db

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Client struct {
	db *gorm.DB
}

func NewClient() (*Client, error) {
	db, err := open("localhost", "root", "root", "5432", "postgres")
	if err != nil {
		return nil, err
	}

	return &Client{
		db: db,
	}, nil
}

func open(host, username, password, port, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		host, username, password, database, port,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func (c *Client) Conn(ctx context.Context) *gorm.DB {
	return c.db.WithContext(ctx)
}
