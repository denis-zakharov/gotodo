package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/denis-zakharov/gotodo/cmd"
	"github.com/denis-zakharov/gotodo/db"
)

func main() {
	h := os.Getenv("HOME")
	dbDir := filepath.Join(h, ".cache/gotodo")
	must(os.MkdirAll(dbDir, 0755))
	must(db.Init(filepath.Join(dbDir, "tasks.db")))

	cmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
