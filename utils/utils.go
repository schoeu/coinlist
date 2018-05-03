package utils

import (
	"database/sql"
	"log"
	"strings"
	"time"
)

// error handler.
func ErrHandle(err error) {
	if err != nil {
		log.Println(err)
	}
}

// db opener
func OpenDb(dbTyepe string, dbStr string) *sql.DB {
	if dbTyepe == "" {
		dbTyepe = "mysql"
	}
	db, err := sql.Open(dbTyepe, dbStr)
	ErrHandle(err)

	err = db.Ping()
	ErrHandle(err)
	return db
}

func GetCurrentDate() string {
	t := time.Now().String()
	return strings.Split(t, " ")[0]
}
