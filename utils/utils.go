package utils

import (
	"log"
	"database/sql"
)

// error handler.
func ErrHandle(err error) {
	if err != nil {
		log.Println(err)
	}
}

// 创建数据库链接
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