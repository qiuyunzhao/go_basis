package main

import (
	_ "github.com/mattn/go-sqlite3"
	"go_code/34_数据库操作/5_sqlite3/model"
)

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
