package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

var (
	Url  = os.Getenv("Url")
	conn *pgx.Conn
	err  error
)

func init() {
	log.Println("Database connecting ..")
	log.Println("database: ", Url)
	conn, err = pgx.Connect(context.Background(), Url)
	if err != nil {
		log.Fatal("Error connecting to db", err)
	}
}
func GetClient() *pgx.Conn {
	return conn
}
