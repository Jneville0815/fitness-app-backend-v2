package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlConnection() (*sqlx.DB, error) {
	maxConn := 100
	maxIdleConn := 10
	maxLifetimeConn := 2

	db, err := sqlx.Connect("mysql", "root:1baseball1@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
