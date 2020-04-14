package main

import (
	_ "github.com/mattn/go-sqlite3"
	"go_basis/34_数据库操作/5_sqlite3/model"
)

// 必看 http://note.youdao.com/noteshare?id=a459d518fd9d284a838d5e37beadba1b
func main() {

	model.InitSQLiteDB()
	SQLite := model.SQLite
	defer SQLite.SQLite3.Close()

	SQLite.Ping()
	SQLite.CreateTable()
	//SQLite.Insert()
	//SQLite.Query()
	//SQLite.Delete()
	//SQLite.Update()
	//SQLite.QueryTables()
	//SQLite.DropTable()
	SQLite.Transaction()
}
