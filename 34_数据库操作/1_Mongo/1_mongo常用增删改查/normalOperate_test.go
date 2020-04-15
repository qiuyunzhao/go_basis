/*
@ Time : 2020/4/15 9:15
@ Author : qyz
@ File : normalOperate_test
@ Software: GoLand
@ Description: 数据库相关操作测试文件
*/

package main

import (
	"encoding/json"
	"log"
	"testing"
)

//----------------------------------- 新增 ------------------------------------
func TestInsert(t *testing.T) {
	err := Insert(user1, user2) //插入数据，一次可以插入多个

	if err != nil {
		log.Println(err)
	}
}

//----------------------------------- 删除 ------------------------------------
//删除时没有找到 err 为 not found
func TestDelete(t *testing.T) {
	//err := removeById("5e966156300118702c7905d6") //根据 _id 删除一条
	//err := removeOneByField()                     //根据属性删除一条
	_, err := removeAllByField() //根据属性删除多条

	if err != nil {
		log.Println(err)
	}
}

//----------------------------------- 修改 ------------------------------------
//update 没有找到 err为not found ；upset 没有找到会新增
func TestUpdate(t *testing.T) {
	//err := updateById()   //根据Id修改
	//_, err := upsetById() //根据Id修改,没有会新增

	//err := UpdateOneByField() //修改满足条件的一条数据
	_, err := UpdateAllByField() //修改满足条件的所有条数据

	if err != nil {
		log.Println(err)
	}
}

//----------------------------------- 查询 ------------------------------------
//没有找到 err 为 not found
func TestQuery(t *testing.T) {
	//err, data := FindAll() //查询全部
	//err, data := findById() //根据 _id 查询
	//err, data := findOneByField() //根据字段查询一条
	err, data := findManyByField() //根据字段查询多条

	//----------------------------------------------------------------------------
	if err != nil {
		log.Println(err)
	}
	result, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(string(result))
	}
}
