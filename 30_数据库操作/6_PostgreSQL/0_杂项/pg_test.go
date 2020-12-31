/*
@ Time : 2020/11/4 13:49
@ Author : qyz
@ File : pgUtils_test
@ Software: GoLand
@ Description:
*/

package __杂项

import (
	"fmt"
	"go_basis/30_数据库操作/6_PostgreSQL/entites"
	"go_basis/30_数据库操作/6_PostgreSQL/pgUtils"
	"log"
	"testing"
)

func Test1(t *testing.T) {
	pgCon := pgUtils.PGConnection{}
	if err := pgCon.ConnDB(); err != nil {
		return
	}
	defer pgCon.CloseDB()
	// 获取到数据库中所有的表，字段，索引的信息
	if tables, err := pgCon.Engine.DBMetas(); err != nil {
		log.Println(err)
	} else {
		fmt.Println(tables)
	}
	// TableInfo()
	//根据传入的结构体指针及其对应的Tag，提取出模型对应的表结构信息。这里不是数据库当前的表结构信息，而是通过struct建模时希望数据库的表的结构信息
	if table, err := pgCon.Engine.TableInfo(new(entites.User)); err != nil {
		log.Println(err)
	} else {
		fmt.Println(table)
	}

}
