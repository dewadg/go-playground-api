package adapter

import (
	"database/sql"
	"os"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var sqliteOnce sync.Once
var sqliteDB *sql.DB

func GetSqliteDB() (*sql.DB, error) {
	var err error
	sqliteOnce.Do(func() {
		sqliteDB, err = ConnectSqlite(os.Getenv("SQLITE_DSN"))
	})

	return sqliteDB, err
}

func ConnectSqlite(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(2)
	db.SetConnMaxLifetime(1 * time.Hour)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
