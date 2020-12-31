/*
@ Time : 2020/11/3 15:36
@ Author : qyz
@ File : add_test
@ Software: GoLand
@ Description:
*/

package add

import (
	"fmt"
	"go_basis/30_数据库操作/6_PostgreSQL/entites"
	"log"
	"testing"
)

func TestAddUser_ORM(t *testing.T) {
	user := entites.User{
		Id:       4,
		UserName: "jerry",
		Age:      8,
		Password: "rrrrrr",
		Phone:    "888888",
		Email:    "jerry@qq.com",
	}
	if affected, err := addUser_ORM(user); err != nil {
		log.Println(err)
	} else {
		fmt.Println("受影响行数 ：", affected)
	}
}

func TestAddUser_SQL(t *testing.T) {
	user := entites.User{
		Id:       5,
		UserName: "jemmey",
		Age:      11,
		Password: "yyyyyy",
		Phone:    "4444444",
		Email:    "jerry@ins.com",
	}
	if res, err := addUser_SQL(user); err != nil {
		log.Println(err)
	} else {
		fmt.Println("受影响行数 ：", res)
	}
}
