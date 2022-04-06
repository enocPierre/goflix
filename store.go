package main

// sqlx
import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type Store interface {
	Open() error
	Close() error
}

// DB
type dbStore struct {
	db *sqlx.DB
}

var schema = `
CREATE TABLE IF NOT EXISTS movie
(
	id INTEGER PRIMARy KEY AUTOINCREMENT,
	title TEXT,
	release_date TEXT,
	duration INTEGER,
	trailer_url TEXT
)
`

// implement interface
func (store *dbStore) Open() error {
	db, err := sqlx.Connect("sqlite3", "goflix.db")
	if err != nil {
		return err
	}

	log.Println("Connected to DB")
	db.MustExec(schema)
	store.db = db
	return nil
}

func (store *dbStore) Close() error {
	return store.db.Close()
}