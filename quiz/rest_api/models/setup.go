package models

import "database/sql"

func ConnectDatabase() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/simple_crud")
	if err != nil {
		return nil
	}
	return db
}
