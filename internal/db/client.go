package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Client struct {
	db *gorm.DB
}

type Env struct {
	Hostname string // GGG_HOST
	Port     string // GGG_PORT
	User     string // GGG_USER
	Password string // GGG_PASSWORD
	Dbname   string // GGG_DBNAME
}

func NewClient() (*Client, error) {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}
	e := &Env{
		Hostname: os.Getenv("GGG_HOSTNAME"),
		Port:     os.Getenv("GGG_PORT"),
		User:     os.Getenv("GGG_USERNAME"),
		Password: os.Getenv("GGG_PASSWORD"),
		Dbname:   os.Getenv("GGG_DATABASE"),
	}
	db, err := open(e.Hostname, e.User, e.Password, e.Port, e.Dbname)
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
