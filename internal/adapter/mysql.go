package adapter

import (
	"database/sql"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlOnce sync.Once
var mysqlDB *sql.DB

func GetMysqlDB() (*sql.DB, error) {
	var err error
	mysqlOnce.Do(func() {
		mysqlDB, err = ConnectMysql(os.Getenv("MYSQL_DSN"))
	})

	return mysqlDB, err
}

func ConnectMysql(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
