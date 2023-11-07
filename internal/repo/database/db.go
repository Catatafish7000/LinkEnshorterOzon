package database

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	Host     = "localhost"
	Port     = 5432
	Username = "el"
	Password = "ozon"
	DBName   = "db"
)

type ConfigDB struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type Cfg struct {
}
type Link struct {
	URL       string
	Hash      string
	CreatedAt time.Time
}
type repo struct {
	DB *sqlx.DB
}

func NewRepo() *repo {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, Username, Password, DBName)
	db, err := sqlx.Open("postgres", psqlConn)
	if err != nil {
		log.Fatal("Failed to connect to db")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Failed to ping db")
	}
	return &repo{DB: db}
}

func (r *repo) SaveHashByURL(ctx context.Context, url, hash string) error {
	//psqlQ := fmt.Sprintf("INSERT INTO %s (url,hash,created_at) VALUES ($1,$2,$3);", URLsTableName)

	_, err := r.DB.ExecContext(ctx, "INSERT INTO urls (longurl,shorturl,created_at) VALUES ($1,$2,$3)", url, hash, time.Now())
	return err
}

func (r *repo) GetURL(ctx context.Context, hash string) (string, error) {
	var url string
	if err := r.DB.QueryRowContext(ctx, "select longurl from urls where shorturl=$1", hash).Scan(&url); err == nil {
		return url, err
	} else {
		return "", err
	}
}

func (r *repo) Clear(ctx context.Context) {
	//current := time.Now()
	_, err := r.DB.ExecContext(ctx, "delete from urls where created_at>(CURRENT_TIMESTAMP - interval '1 day';")
	if err != nil {
		log.Println(fmt.Sprintf("Failed to clear memory. Error: %v", err))
	}
}
