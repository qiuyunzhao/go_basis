/*
@ Time : 2020/11/3 15:45
@ Author : qyz
@ File : User
@ Software: GoLand
@ Description:
*/

package entites

import (
	"go_basis/28_数据库操作/6_PostgreSQL/pgUtils"
)

//用户表结构
type User struct {
	Id       int    `xorm:"int(32) notnull unique"`
	UserName string `xorm:"varchar(255)"`          // 支持驼峰命名法自动映射
	Age      int    `xorm:"int"`                   // 名称一致的自动映射
	Password string `xorm:"varchar(255)"`          // 名称一致的自动映射
	Phone    string `xorm:"varchar(255)"`          // 名称一致的自动映射
	Email    string `xorm:"varchar(255) 'e_mail'"` // 指定映射数据库password字段
}

func (user *User) queryUserById(id int) error {
	pgConn := pgUtils.PGConnection{}
	if err := pgConn.ConnDB(); err != nil {
		return err
	}
	defer pgConn.CloseDB()

	_, err := pgConn.Engine.Table("user").Where("id=?", id).Get(user)
	if err != nil {
		return err
	}
	return nil
}
