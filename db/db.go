package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// ~ ok so main thing that this particular driver is doing is registering the driver and preparing the connection pool manager for handling the connections
func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		return nil, err
	}

	return db, nil
}
