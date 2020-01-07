package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

// 检查是否有错
func checkErr(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
	} else {
		fmt.Println(msg, "OK")
	}
}

//-------------------------------------------------------- Ping --------------------------------------------------------
func (s *SQLiteDB) Ping() {
	err := s.SQLite3.Ping()

	checkErr("Ping : ", err)
}

// --------------------------------------------------- 删除数据库文件 ---------------------------------------------------
func removeDBfile(dbFilePath string) {
	err := os.Remove(dbFilePath)
	if err != nil {
		log.Println("删除数据库文件：", err)
	}
}

//------------------------------------------------------- 创建表 -------------------------------------------------------
func (s *SQLiteDB) CreateTable() {
	sql := s.SQLite3

	//sqlStmt := `create table userinfo (username text , departname text , created text);`
	sqlStmt := `CREATE TABLE if NOT EXISTS userinfo (
                       id INTEGER PRIMARY KEY	AUTOINCREMENT	NOT NULL ,
                       username   TEXT	NOT NULL,
                       departname TEXT,
                       created    TEXT
                       );`

	_, err := sql.Exec(sqlStmt)

	checkErr("Create Table : ", err)
}

//-------------------------------------------------- 查询数据库中的表 ---------------------------------------------------
func (s *SQLiteDB) QueryTables() {
	sql := s.SQLite3

	rows, err := sql.Query("SELECT name FROM sqlite_master WHERE type='table' order by name; ")
	checkErr("Query", err)

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		fmt.Println("表名：", tableName)
	}
}

// -------------------------------------------------- 删除数据库中的表 --------------------------------------------------
func (s *SQLiteDB) DropTable() {
	sql := s.SQLite3

	res, err := sql.Exec("DROP TABLE userinfo; ")

	checkErr("DropTables", err)
	fmt.Println(res)
}

// ------------------------------------------------------ 插入数据 -----------------------------------------------------
func (s *SQLiteDB) Insert() {
	sql := s.SQLite3

	stmt, err := sql.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	res, err := stmt.Exec("邱浩骞", "IEC118", time.Now())
	id, err := res.LastInsertId()

	checkErr("Insert : ", err)
	fmt.Println("LastInsertId", id)
}

// ------------------------------------------------------ 查询数据 ------------------------------------------------------
func (s *SQLiteDB) Query() {
	sql := s.SQLite3

	rows, err := sql.Query("SELECT * FROM userinfo")

	checkErr("Query : ", err)

	for rows.Next() {
		var id int64
		var username string
		var department string
		var created string
		err = rows.Scan(&id, &username, &department, &created)
		fmt.Println(id, "--", username, "--", department, "--", created)
	}
}

// ------------------------------------------------------ 删除数据 -----------------------------------------------------
func (s *SQLiteDB) Delete() {
	sql := s.SQLite3

	stmt, err := sql.Prepare("DELETE  FROM userinfo WHERE rowid = ?")
	result, err := stmt.Exec(1)
	affectNum, err := result.RowsAffected()

	checkErr("Delete : ", err)
	fmt.Println("delete affect rows is ", affectNum)
}

// ------------------------------------------------------ 数据更新 ------------------------------------------------------
func (s *SQLiteDB) Update() {
	sql := s.SQLite3

	stmt, err := sql.Prepare("UPDATE userinfo SET created = ? WHERE rowid = ?")
	result, err := stmt.Exec("2016-09-7", 2)
	affectNum, err := result.RowsAffected()

	checkErr("Update", err)
	fmt.Println("update affect rows is ", affectNum)
}

// ------------------------------------------------------- 事务 --------------------------------------------------------
func (s *SQLiteDB) Transaction() {
	sql := s.SQLite3

	tx, err := sql.Begin()
	if err != nil {
		log.Println("开启SQLite事务失败")
		return
	}
	defer clearTransaction(tx) //如果执行过程中都没有 commit和roolback那就由这个来收尾

	//1、新增
	stmt1, err := tx.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	res, err := stmt1.Exec("邱浩骞", "IEC118", time.Now())
	if err != nil {
		log.Println("新增失败，执行回滚", err)
		tx.Rollback()
		return
	}

	//获取新增数据的id
	id, err := res.LastInsertId()
	if err != nil {
		log.Println("新增数据id获取失败，执行回滚", err)
		tx.Rollback()
		return
	}
	fmt.Println("新增数据id: ", id) //这里需要的话，可以改成判断函数，等于小于0，rollback

	//2、修改
	stmt2, err := tx.Prepare("UPDATE userinfo SET username=?,created = ? WHERE id = ?")
	res2, err := stmt2.Exec("小明", "2016-09-7", id)
	if err != nil {
		log.Println("更新失败，执行回滚：", err)
		tx.Rollback()
		return
	}

	affectNum, err := res2.RowsAffected()
	if err != nil {
		log.Println("更新数据获取受影响条数失败，执行回滚", err)
		tx.Rollback()
		return
	}
	fmt.Println("更新数据受影响提条数: ", affectNum) //这里需要的话，可以改成判断函数，等于小于0，rollback

	err = tx.Commit()
	if err != nil {
		log.Println("提交事务失败：", err)
	}
}

//事务回滚
func clearTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	fmt.Println("defer 执行了rollback")
	if err != sql.ErrTxDone && err != nil {
		log.Println(err)
	}
}
