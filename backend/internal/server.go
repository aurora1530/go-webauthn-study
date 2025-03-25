package internal

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Server struct {
	*sql.DB
}

func openDB() *sql.DB {
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := "db"
	port := "5432"

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
	return db
}

func NewServer() Server{
	db := openDB()
	return Server{db}
}
