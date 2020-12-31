/*
@ Time : 2020/4/15 9:47
@ Author : qyz
@ File : aggregateOperate_test.go
@ Software: GoLand
@ Description:
*/

package main

import (
	"encoding/json"
	"log"
	"testing"
)

//--------------------------------------------------- $project -------------------------------------------------------
//$project 修改输入文档的结构。可以用来重命名、增加或删除域，也可以用于创建计算结果以及嵌套文档。（SELECT）
func TestProject(t *testing.T) {
	err, result := projectPipe()

	//将对象，转换成json格式展示
	data, err := json.Marshal(result)
	if err != nil {
		log.Println("err:\t", err.Error())
	}
	log.Println(string(data))
}

//--------------------------------------------------- $match -------------------------------------------------------
//$match 用于过滤数据，只输出符合条件的文档。（WHERE、HAVING）
func TestMatch(t *testing.T) {
	err, result := matchPipe()

	//将对象，转换成json格式展示
	data, err := json.Marshal(result)
	if err != nil {
		log.Println("err:\t", err.Error())
	}
	log.Println(string(data))
}

//--------------------------------------------------- $sort --------------------------------------------------------
//$sort 将输入文档排序（ORDER BY）
func TestSort(t *testing.T) {
	err, result := sortPipe()

	//将对象，转换成json格式展示
	data, err := json.Marshal(result)
	if err != nil {
		log.Println("err:\t", err.Error())
	}
	log.Println(string(data))
}

//--------------------------------------------------- $limit -------------------------------------------------------
//$limit 用来限制MongoDB聚合管道返回的文档数（LIMIT）
func TestLimit(t *testing.T) {
	err, result := limitPipe()

	//将对象，转换成json格式展示
	data, err := json.Marshal(result)
	if err != nil {
		log.Println("err:\t", err.Error())
	}
	log.Println(string(data))
}

//--------------------------------------------------- 比较管道 ------------------------------------------------------
//$eq、$ne、$gt、$lt、"$gte"、$lte
// $eq：等于         $ne：不等于
// $gt：大于         $lt：小于
// $gte：大于等于    $lte：小于等于 （可用于判断字符串）
func TestCompare(t *testing.T) {
	err, result := comparePipe()

	//将对象，转换成json格式展示
	data, err := json.Marshal(result)
	if err != nil {
		log.Println("err:\t", err.Error())
	}
	log.Println(string(data))
}

//--------------------------------------------------- 或运算 ------------------------------------------------------
func TestOrPipe(t *testing.T) {
	err, result := OrPipe()

	//将对象，转换成json格式展示
	data, err := json.Marshal(result)
	if err != nil {
		log.Println("err:\t", err.Error())
	}
	log.Println(string(data))
}

//--------------------------------------------------- $skip管道 ----------------------------------------------------
//$skip 在聚合管道中跳过指定数量的文档，并返回余下的文档(效率低,慎用)
func TestSkip(t *testing.T) {
	err, result := skipPipe()

	//将对象，转换成json格式展示
	data, err := json.Marshal(result)
	if err != nil {
		log.Println("err:\t", err.Error())
	}
	log.Println(string(data))
}

//--------------------------------------------------- $group管道 ---------------------------------------------------
// $group 将集合中的文档分组，可用于统计结果（GROUP BY）(未研究透彻)
// $abs、$max、$min、$avg、$sum、$first、$last、$unwind、$addToSet、$push 配合使用
func TestGroup(t *testing.T) {
	err, result := groupPipe()

	//将对象，转换成json格式展示
	data, err := json.Marshal(result)
	if err != nil {
		log.Println("err:\t", err.Error())
	}
	log.Println(string(data))
}

//--------------------------------------------------- $lookup管道 --------------------------------------------------
// $lookup 用以引入其它集合的数据，表关联查询（join）
func TestLookup(t *testing.T) {
	err, result := lookupPipe()

	//将对象，转换成json格式展示
	data, err := json.Marshal(result)
	if err != nil {
		log.Println("err:\t", err.Error())
	}
	log.Println(string(data))
}
