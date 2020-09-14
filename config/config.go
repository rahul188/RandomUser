package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//import db
func GetDB() (db *sql.DB, err error) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "db_random"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
