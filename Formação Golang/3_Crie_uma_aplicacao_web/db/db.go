package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaBancoDeDados() *sql.DB {
	connectionString := "user=postgres password=root dbname=alura_loja host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return db
}
