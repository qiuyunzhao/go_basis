package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//参考： https://www.jianshu.com/p/f1d4300c0067
const (
	url       = "127.0.0.1:27017"
	database1 = "AggregationPipelineTest"
)

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
	//-------------------- mongoDB聚合管道(多个管道表达式则在上一个表达式基础上进行下一个表达式的筛选)----------------------
	session, err := mgo.Dial(url)
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	session.SetMode(mgo.Monotonic, true)
	//切换到数据库
	database := session.DB(database1)

	//---------------$project 修改输入文档的结构。可以用来重命名、增加或删除域，也可以用于创建计算结果以及嵌套文档。（SELECT）
	//var commodity []Commodity
	//c := database.C("commodity") //查询哪张表
	//
	//err = c.Pipe([]bson.M{
	//	{"$project": bson.M{"name": 1, "price": 1}}, //将要查询的字段赋值为 1 (_id默认会查询)
	//}).All(&commodity)
	//
	////将对象，转换成json格式展示
	//data, err := json.Marshal(commodity)

	//--------------------------------------------------------$match 用于过滤数据，只输出符合条件的文档。（WHERE、HAVING）
	//var commodity []Commodity
	//c := database.C("commodity") //查询哪张表
	//
	//err = c.Pipe([]bson.M{
	//	{"$match": bson.M{"name": "苹果电脑"}},
	//}).All(&commodity)
	//
	////将对象，转换成json格式展示
	//data, err := json.Marshal(commodity)

	// ----------------------------------------------------------------------------------$sort 将输入文档排序（ORDER BY）
	//var commodity []Commodity
	//c := database.C("commodity") //查询哪张表
	//
	//err = c.Pipe([]bson.M{
	//	{"$sort": bson.M{"price": 1}}, // -1：降序排序 1：升序排序
	//}).All(&commodity)
	//
	////将对象，转换成json格式展示
	//data, err := json.Marshal(commodity)

	// ----------------------------------------------------------------$limit 用来限制MongoDB聚合管道返回的文档数（LIMIT）
	//var commodity []Commodity
	//c := database.C("commodity") //查询哪张表
	//
	//err = c.Pipe([]bson.M{
	//	{"$match": bson.M{"price": bson.M{"$gte": 1999}}}, //价格>=1999
	//	{"$limit": 3}, //限制三条
	//}).All(&commodity)
	//
	////将对象，转换成json格式展示
	//data, err := json.Marshal(commodity)

	// --------------------------------------------------$skip 在聚合管道中跳过指定数量的文档，并返回余下的文档(效率低,慎用)
	//var commodity []Commodity
	//c := database.C("commodity") //查询哪张表
	//
	//err = c.Pipe([]bson.M{
	//	{"$match": bson.M{"price": bson.M{"$gte": 1999}}}, //价格>=1999
	//	{"$sort": bson.M{"price": -1}}, // -1：降序排序 1：升序排序
	//	{"$skip": 2}, //限制三条
	//}).All(&commodity)
	//
	////将对象，转换成json格式展示
	//data, err := json.Marshal(commodity)

	// ---------------------------------------------------$group 将集合中的文档分组，可用于统计结果（GROUP BY）(未研究透彻)
	// ------------------------------------------- $abs、$max、$min、$avg、$sum、$first、$last、$unwind、$addToSet、$push
	var groupBy []interface{}
	c := database.C("commodity") //查询哪张表

	err = c.Pipe([]bson.M{
		//_id:合并条件的字段（好像合并条件必须用 _id 接收）
		//{"$group": bson.M{"_id": "$manufacturer", "num": bson.M{"$sum": 1}}},//按厂商分组，统计商品种类
		//{"$group": bson.M{"_id": "$manufacturer", "num": bson.M{"$sum": "$count"}}}, //按厂商分组，统计各类产品总数量
		//{"$group": bson.M{"_id": "$manufacturer", "avgPrice": bson.M{"$avg": "$price"}}}, //按厂商分组，统计各种商品种类的平均价格
		//{"$group": bson.M{"_id": "$manufacturer", "maxPrice": bson.M{"$max": "$price"}}}, //按厂商分组,找出最贵产品的价格
		//{"$group": bson.M{"_id": "$manufacturer", "minPrice": bson.M{"$min": "$price"}}}, //按厂商分组,找出最便宜产品价格
		//{"$group": bson.M{"_id": "$manufacturer", "first": bson.M{"$first": "$price"}}}, //按厂商分组,找出第一条产品价格
		//{"$group": bson.M{"_id": "$manufacturer", "last": bson.M{"$last": "$price"}}}, //按厂商分组,找出最后一条产品价格
		//{"$group": bson.M{"_id": "$manufacturer", "last": bson.M{"$last": "$price"}}}, //按厂商分组,找出最后一条产品价格

		//$addToSet 将分组后的某一个字段放到一个数组中，重复的元素只出现一次，加入无顺序
		//{"$group": bson.M{"_id": "$price", "name": bson.M{"$addToSet": "$name"}}}, //按价格分组,将商品名称放入一个数组（不重复）
		//{"$group": bson.M{"_id": "$price", "name": bson.M{"$push": "$name"}}}, //按价格分组,将商品名称放入一个数组（可重复）

		//$unwind用来实现对文档的拆分,可以将文档中的值拆分为单独的文档
		//{"$match": bson.M{"name": "鼠标"}},
		//{"$unwind":"$class"},
	}).All(&groupBy)

	//将对象，转换成json格式展示
	data, err := json.Marshal(groupBy)

	// ------------------------------------------------------------------$lookup 用以引入其它集合的数据，表关联查询（join）
	//var article []interface{}
	//c := database.C("article") //查询表（联查数据在该表的字段上进行扩展）
	//err = c.Pipe([]bson.M{
	//	// "from": "联查的表名", "localField": "查询表外键", "foreignField": "联查对应外键", "as": "联查数据在查询数据中的字段名"
	//	{"$lookup": bson.M{"from": "author", "localField": "authorId", "foreignField": "authorId", "as": "author"}},
	//	{"$lookup": bson.M{"from": "articlecate", "localField": "cateId", "foreignField": "cateId", "as": "articlecate"}},
	//}).All(&article)
	//
	////将对象，转换成json格式展示
	//data, err := json.Marshal(article)

	// -------------------------------------------------------------------------------- $eq、$ne、$gt、$lt、"$gte"、$lte
	//// $eq：等于     $ne：不等于     $gt：大于     $lt：小于     $gte：大于等于     $lte：小于等于 （可用于判断字符串）

	//var commodity []Commodity
	////var commodity []interface{}
	//c := database.C("commodity") //查询哪张表
	//
	//pipLine := []bson.M{
	//	{"$match": bson.M{"price": bson.M{"$gte": 1999},}}, //价格>=1999
	//	//{"$match": bson.M{"manufacturer": bson.M{"$eq": "小米"}}}, //制造商为“小米”
	//	{"$match": bson.M{"manufacturer": "小米"}}, //制造商为“小米”
	//	{"$sort": bson.M{"price": 1}}, // -1：降序排序 1：升序排序
	//}
	//
	//err = c.Pipe(pipLine).All(&commodity)
	//
	////将对象，转换成json格式展示
	//data, err := json.Marshal(commodity)

	//------------------------------------------------------------------------------------------------------------------
	if err != nil {
		fmt.Println("err:\t", err.Error())
	}
	fmt.Println(string(data))
	fmt.Println("------")

	if err != nil {
		log.Println(err)
	}

}
