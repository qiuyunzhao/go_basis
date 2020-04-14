package main

import (
	"encoding/json"
	"fmt"
	"go_basis/34_数据库操作/1_Mongo/mongoUtils"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// 聚合管道用于复杂查询
// mongoDB聚合管道(多个管道表达式则在上一个表达式基础上进行下一个表达式的筛选)
// 参考： https://www.jianshu.com/p/f1d4300c0067

const databaseName = "AggregationPipelineTest"

//商品
type Commodity struct {
	Id           bson.ObjectId `bson:"_id"`
	Name         string        `bson:"name"`
	Manufacturer string        `bson:"manufacturer"`
	Price        int           `bson:"price"`
	Count        int           `bson:"count"`
}

//文章
type Article struct {
	Id          bson.ObjectId `bson:"_id"`
	CateId      string        `bson:"cateId"`
	AuthorId    string        `bson:"authorId"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	Time        string        `bson:"time"`
}

//文章分类
type Articlecate struct {
	Id          bson.ObjectId `bson:"_id"`
	CateId      string        `bson:"cateId"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
}

//文章作者
type Author struct {
	Id         bson.ObjectId `bson:"_id"`
	AuthorId   string        `bson:"authorId"`
	AuthorName string        `bson:"authorName"`
	PassWord   string        `bson:"passWord"`
	Age        string        `bson:"age"`
	Sex        string        `bson:"sex"`
	Phone      string        `bson:"phone"`
}

func main() {
	//--------------------------------------------------- $project -----------------------------------------------------
	//err, result := projectPipe()
	//--------------------------------------------------- $match -------------------------------------------------------
	//err, result := matchPipe()
	//--------------------------------------------------- $sort --------------------------------------------------------
	//err, result := sortPipe()
	//--------------------------------------------------- $limit -------------------------------------------------------
	//err, result := limitPipe()
	//--------------------------------------------------- 比较管道 ------------------------------------------------------
	//err, result := comparePipe()
	//--------------------------------------------------- $skip管道 ----------------------------------------------------
	//err, result := skipPipe()
	//--------------------------------------------------- $group管道 ---------------------------------------------------
	//err, result := groupPipe()
	//--------------------------------------------------- $lookup管道 --------------------------------------------------
	err, result := lookupPipe()

	//------------------------------------------------------------------------------------------------------------------
	//将对象，转换成json格式展示
	data, err := json.Marshal(result)
	if err != nil {
		log.Println("err:\t", err.Error())
	}
	fmt.Println(string(data))
}

//======================================================================================================================
//======================================================================================================================

//------------------ $project 修改输入文档的结构。可以用来重命名、增加或删除域，也可以用于创建计算结果以及嵌套文档。（SELECT）

func projectPipe() (error, []Commodity) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: "commodity"}
	err := db.ConnDB()
	defer db.CloseDB()
	var commoditys []Commodity

	//指定查询字段，将要查询的字段赋值为 1 (_id默认会查询)
	pipe := []bson.M{
		{"$project": bson.M{"name": 1, "price": 1}},
	}

	err = db.Collection.Pipe(pipe).All(&commoditys)
	return err, commoditys
}

//------------------------------------------------------------ $match 用于过滤数据，只输出符合条件的文档。（WHERE、HAVING）

func matchPipe() (error, []Commodity) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: "commodity"}
	err := db.ConnDB()
	defer db.CloseDB()
	var commoditys []Commodity

	// 匹配 name=苹果电脑 的数据
	pipe := []bson.M{
		{"$match": bson.M{"name": "苹果电脑"}},
	}

	err = db.Collection.Pipe(pipe).All(&commoditys)
	return err, commoditys
}

// ---------------------------------------------------------------------------------------$sort 将输入文档排序（ORDER BY）

func sortPipe() (error, []Commodity) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: "commodity"}
	err := db.ConnDB()
	defer db.CloseDB()
	var commoditys []Commodity

	//查询结果进行排序 -1：降序排序 1：升序排序
	pipe := []bson.M{
		{"$sort": bson.M{"price": 1}}, //价格升序
	}

	err = db.Collection.Pipe(pipe).All(&commoditys)
	return err, commoditys
}

// -------------------------------------------------------------------- $limit 用来限制MongoDB聚合管道返回的文档数（LIMIT）

func limitPipe() (error, []Commodity) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: "commodity"}
	err := db.ConnDB()
	defer db.CloseDB()
	var commoditys []Commodity

	//限制查询条数
	pipe := []bson.M{
		{"$match": bson.M{"price": bson.M{"$gte": 1999}}}, //价格>=1999
		{"$sort": bson.M{"price": 1}},                     //价格升序
		{"$limit": 3},                                     //限制三条
	}

	err = db.Collection.Pipe(pipe).All(&commoditys)
	return err, commoditys
}

// ------------------------------------------------------------------------------------ $eq、$ne、$gt、$lt、"$gte"、$lte
// $eq：等于         $ne：不等于
// $gt：大于         $lt：小于
// $gte：大于等于    $lte：小于等于 （可用于判断字符串）
func comparePipe() (error, []Commodity) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: "commodity"}
	err := db.ConnDB()
	defer db.CloseDB()
	var commoditys []Commodity

	pipe := []bson.M{
		{"$match": bson.M{"price": bson.M{"$gte": 1999}}},       //价格>=1999
		{"$match": bson.M{"manufacturer": bson.M{"$eq": "小米"}}}, //制造商为“小米”
		//{"$match": bson.M{"manufacturer": "小米"}}, //制造商为“小米”
		{"$sort": bson.M{"price": 1}}, // -1：降序排序 1：升序排序
	}

	err = db.Collection.Pipe(pipe).All(&commoditys)
	return err, commoditys
}

// ---------------------------------------------------- $skip 在聚合管道中跳过指定数量的文档，并返回余下的文档(效率低,慎用)
func skipPipe() (error, []Commodity) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: "commodity"}
	err := db.ConnDB()
	defer db.CloseDB()
	var commoditys []Commodity

	pipe := []bson.M{
		{"$match": bson.M{"price": bson.M{"$gte": 1999}}}, //价格>=1999
		{"$sort": bson.M{"price": 1}},                     // -1：降序排序 1：升序排序
		{"$skip": 2},                                      //跳过前两条
	}

	err = db.Collection.Pipe(pipe).All(&commoditys)
	return err, commoditys
}

// ---------------------------------------------------$group 将集合中的文档分组，可用于统计结果（GROUP BY）(未研究透彻)
// ------------------------------------- $abs、$max、$min、$avg、$sum、$first、$last、$unwind、$addToSet、$push 配合使用
func groupPipe() (error, []interface{}) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: "commodity"}
	err := db.ConnDB()
	defer db.CloseDB()
	var data []interface{}

	pipe := []bson.M{
		//_id:合并条件的字段（好像合并条件必须用 _id 接收）
		//{"$group": bson.M{"_id": "$manufacturer", "num": bson.M{"$sum": 1}}}, //按厂商分组，统计商品种类
		//{"$group": bson.M{"_id": "$manufacturer", "num": bson.M{"$sum": "$count"}}}, //按厂商分组，统计各类产品总数量
		//{"$group": bson.M{"_id": "$manufacturer", "avgPrice": bson.M{"$avg": "$price"}}}, //按厂商分组，统计各种商品种类的平均价格
		//{"$group": bson.M{"_id": "$manufacturer", "maxPrice": bson.M{"$max": "$price"}}}, //按厂商分组,找出最贵产品的价格
		//{"$group": bson.M{"_id": "$manufacturer", "minPrice": bson.M{"$min": "$price"}}}, //按厂商分组,找出最便宜产品价格
		//{"$group": bson.M{"_id": "$manufacturer", "first": bson.M{"$first": "$price"}}}, //按厂商分组,找出第一条产品价格
		//{"$group": bson.M{"_id": "$manufacturer", "last": bson.M{"$last": "$price"}}}, //按厂商分组,找出最后一条产品价格
		//{"$group": bson.M{"_id": "$manufacturer", "last": bson.M{"$last": "$price"}}}, //按厂商分组,找出最后一条产品价格
		//
		////$addToSet 将分组后的某一个字段放到一个数组中，重复的元素只出现一次，加入无顺序
		//{"$group": bson.M{"_id": "$price", "name": bson.M{"$addToSet": "$name"}}}, //按价格分组,将商品名称放入一个数组（不重复）
		//{"$group": bson.M{"_id": "$price", "name": bson.M{"$push": "$name"}}}, //按价格分组,将商品名称放入一个数组（可重复）
		//
		////$unwind用来实现对文档的拆分,可以将文档中的值拆分为单独的文档
		{"$match": bson.M{"name": "鼠标"}},
		{"$unwind": "$class"},
	}

	err = db.Collection.Pipe(pipe).All(&data)
	return err, data
}

// ------------------------------------------------------------------ $lookup 用以引入其它集合的数据，表关联查询（join）
func lookupPipe() (error, []interface{}) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: "article"}
	err := db.ConnDB()
	defer db.CloseDB()
	var data []interface{}

	pipe := []bson.M{
		// "from": "联查的表名", "localField": "查询表外键", "foreignField": "联查对应外键", "as": "联查数据在查询数据中的字段名"
		{"$lookup": bson.M{"from": "author", "localField": "authorId", "foreignField": "authorId", "as": "author"}},
		{"$lookup": bson.M{"from": "articlecate", "localField": "cateId", "foreignField": "cateId", "as": "articlecate"}},
	}

	err = db.Collection.Pipe(pipe).All(&data)
	return err, data
}
