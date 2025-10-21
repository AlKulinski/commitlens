package main

import (
	"log"

	"github.com/alkowskey/commitlens/cmd"
	"github.com/alkowskey/commitlens/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	version := "1.0.0"
	db := db.InitDb()
	defer db.Close()
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute(version, db)
}
