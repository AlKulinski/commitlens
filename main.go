package main

import (
	"github.com/alkowskey/commitlens/cmd"
	"github.com/alkowskey/commitlens/internal/db"
)

func main() {
	version := "1.0.0"
	db := db.InitDb()
	defer db.Close()

	cmd.Execute(version, db)
}
