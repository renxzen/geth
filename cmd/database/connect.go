package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var Client *sql.DB

func init() {
	godotenv.Load()
	url := os.Getenv("DATABASE_NAME")
	token := os.Getenv("DATABASE_TOKEN")
	uri := fmt.Sprintf("libsql://%s.turso.io?authToken=%s", url, token)

	var err error
	Client, err = sql.Open("libsql", uri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", uri, err)
		return
	}
}
