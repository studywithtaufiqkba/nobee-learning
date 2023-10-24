package database

import (
	"day_6/config"
	"github.com/jmoiron/sqlx"
)

func ConnectingSQLX(cfg config.DB) (db *sqlx.DB, err error) {
	dsn := ""

	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return
	}

	return db, nil
}
