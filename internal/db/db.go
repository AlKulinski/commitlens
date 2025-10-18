package db

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func InitDb() *sql.DB {
	db, err := sql.Open("sqlite", "./internal/db/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	runMigrations(db)
	fmt.Println("Connected to SQLite!")

	return db
}

func runMigrations(db *sql.DB) {

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	log.Println("database migrated sucessfuly")
}
