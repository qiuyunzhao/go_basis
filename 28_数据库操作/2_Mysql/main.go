package main

/**
*golang中mysql的用法
 */

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	ID        int64          `db:"id"`
	Name      sql.NullString `db:"name"`    //数据表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
	Age       sql.NullInt64  `db:"age"`     //实际使用时可以结构体中正常使用基本类型，在查询执行  rows.Scan(&s) 前定义变量 var s sql.NullString 和结构体 user := new(User)
	Address   sql.NullString `db:"address"` //执行rows.Scan(&s) 后 user.Name = s.String
	CreatTime time.Time      `db:"creatTime"`
}

const (
	USERNAME = "root"
	PASSWORD = "zhao17615110273"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "go_test"
)

func main() {

	DB := GetDB()
	defer DB.Close()

	//queryOne(DB)
	queryMulti(DB)
	//insertData(DB)
	//updateData(DB)
	//deleteData(DB)
	//Trans(DB)
}

//--------------------------------------------- 获取sql.DB对象 ---------------------------------------------------------
func GetDB() *sql.DB {
	//charset:编码集   loc：时区   parseTime：对time.Time时间格式的支持（若不设置时间类型无法直接读写数据库）
	//dsn := "root:zhao17615110273@tcp(localhost:3306)/gotest?charset=utf8&loc=Asia%2FShanghai&parseTime=true" //全dsn
	//dsn := "root:zhao17615110273@/gotest" //省略写法
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	dsn += "?charset=utf8&loc=Asia%2FShanghai&parseTime=true"

	db, err := sql.Open("mysql", dsn)
	db.Ping() //正式建立连接，不写不会影响程序执行，只是在后边正式建立连接
	if err != nil {
		panic("Open mysql failed,err: " + err.Error())
	}

	//可选项
	db.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	db.SetMaxOpenConns(100)                  //设置最大连接数
	db.SetMaxIdleConns(16)                   //设置闲置连接数

	return db
}

//------------------------------------------------ 查询单行 ------------------------------------------------------------
func queryOne(db *sql.DB) {
	user := new(User)
	row := db.QueryRow("select * from userinfo where id=?", 1)
	//1：row.scan中的字段必须是按照数据库查询到字段的顺序 一一对应，否则报错
	//2：row必须scan，不然会导致连接无法关闭，会一直占用连接，直到超过设置的生命周期
	if err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.CreatTime); err != nil {
		panic("scan failed, err: " + err.Error())
	}
	fmt.Println(*user)
}

//------------------------------------------------ 查询多行 ------------------------------------------------------------
func queryMulti(db *sql.DB) {
	user := new(User)
	rows, err := db.Query("select * from userinfo where id >= ?", 1)
	if err != nil {
		panic("Query failed,err: " + err.Error())
	}
	//函数执行完后执行func()
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未Scan连接一直占用
		}
	}()
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.CreatTime) //不scan会导致连接不释放
		if err != nil {
			panic("Scan failed,err: " + err.Error())
		}
		fmt.Println(*user)
		fmt.Println(user.Name.String)
		fmt.Println(user.Name.Valid)
	}

}

//------------------------------------------------ 插入数据 ------------------------------------------------------------
func insertData(db *sql.DB) {
	result, err := db.Exec("insert INTO userinfo(name,age,creatTime) values(?,?,?)", "qq", 24, time.Now())
	if err != nil {
		panic("Insert failed,err: " + err.Error())
	}

	lastInsertID, err := result.LastInsertId() //插入数据的主键id
	if err != nil {
		panic("Get lastInsertID failed,err: " + err.Error())
	}
	fmt.Println("LastInsertID:", lastInsertID)

	rowsaffected, err := result.RowsAffected() //受影响的行数
	if err != nil {
		panic("Get RowsAffected failed,err: " + err.Error())
	}
	fmt.Println("RowsAffected:", rowsaffected)
}

//------------------------------------------------ 更新数据 ------------------------------------------------------------
func updateData(db *sql.DB) {
	result, err := db.Exec("UPDATE userinfo set age=? where id=?", "30", 3)
	if err != nil {
		panic("update failed,err: " + err.Error())
	}

	//更新数据不返回LastInsertID
	rowsaffected, err := result.RowsAffected() //受影响的行数（更新的内容与原来相同 rowsaffected=0）
	if err != nil {
		panic("Get RowsAffected failed,err: " + err.Error())
	}
	fmt.Println("RowsAffected:", rowsaffected)
}

//------------------------------------------------ 删除数据 ------------------------------------------------------------
func deleteData(db *sql.DB) {
	result, err := db.Exec("delete from userinfo where id=?", 4)
	if err != nil {
		panic("delete failed,err: " + err.Error())
	}

	//更新数据不返回LastInsertID
	rowsaffected, err := result.RowsAffected() //受影响的行数(删除的数据不存在 rowsaffected=0)
	if err != nil {
		panic("Get RowsAffected failed,err: " + err.Error())
	}
	fmt.Println("RowsAffected:", rowsaffected)
}

//-------------------------------------------------- 事务 --------------------------------------------------------------
func Trans(db *sql.DB) {
	tx, err := db.Begin() //开启事务
	if err != nil {
		panic("db.Begin transent error: " + err.Error())
	}

	isCommit := true
	defer func() {
		if isCommit {
			tx.Commit() //提交事务
			fmt.Println("commit...")
		} else {
			tx.Rollback() //事务回滚
			fmt.Println("Rollback...")
		}
	}()

	_, err = tx.Exec("insert INTO userinfo(name,age,creatTime) values(?,?,?)", "transent1", 1, time.Now())
	if err != nil {
		isCommit = false
	}
	_, err = tx.Exec("insert INTO userinfo(name,age,creatTime) values(?,?,?)", "transent2", 2, "错误")
	if err != nil {
		isCommit = false
	}
	_, err = tx.Exec("insert INTO userinfo(name,age,creatTime) values(?,?,?)", "transent3", 3, "2019-03-21")
	if err != nil {
		isCommit = false
	}
}
