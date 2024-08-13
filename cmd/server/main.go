package main

import (
	"ddcard/internal/database"
	"ddcard/internal/server"
	"ddcard/internal/utils"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {

	token, _ := utils.ReadEnvVar("TOKEN")
	db_url, err := utils.ReadEnvVar("TURSO")

	if err != nil {
		log.Println(err)
		db_url = "file:testing.db"
	}

	database.DB = sqlx.MustConnect("libsql", db_url+token)
	server := server.NewServer()

	err = server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
