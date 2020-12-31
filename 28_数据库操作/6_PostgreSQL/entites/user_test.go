/*
@ Time : 2020/11/3 16:18
@ Author : qyz
@ File : user_test
@ Software: GoLand
@ Description:
*/

package entites

import (
	"fmt"
	"testing"
)

func TestQueryUserById(t *testing.T) {
	user1 := User{}
	err := user1.queryUserById(1)
	if err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println("结果：", user1)
	}

	user2 := User{}
	err = user2.queryUserById(2)
	if err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println("结果：", user2)
	}
}
