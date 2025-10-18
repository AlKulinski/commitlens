package main

import (
	"github.com/alkowskey/commit-suggester/cmd"
	"github.com/alkowskey/commit-suggester/internal/db"
)

func main() {
	version := "1.0.0"
	cmd.Execute(version)
	db := db.InitDb()
	defer db.Close()
}
