package model

import (
	"database/sql"
	"log"
)

// 外部利用InitDB()对其进行连接后，可反复使用
var SQLite *SQLiteDB

type SQLiteDB struct {
	SQLite3 *sql.DB
}

// 初始化 SQLite
func InitSQLiteDB() {
	SQLite = &SQLiteDB{
		SQLite3: connectDB(),
	}
}

func connectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "foo.db")
	if err != nil {
		log.Println("sqlite3连接错误!!!")
	}
	return db
}
