package main

import (
	"encoding/json"
	"fmt"
	"go_code/34_数据库操作/1_Mongo/mongoUtils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type User struct {
	Id        bson.ObjectId `bson:"_id"`
	Username  string        `bson:"name"`
	Password  string        `bson:"pass"`
	Age       int64         `bson:"age"`
	Interests []string      `bson:"interests"`
}

const (
	databaseName   = "userinfo" //数据库名称
	collectionName = "user"     //表名称
)

var user1 = User{
	Id:        bson.NewObjectId(),
	Username:  "陈钰琪",
	Password:  "88888888",
	Age:       24,
	Interests: []string{"唱歌", "演戏", "跳舞"},
}

var user2 = User{
	Id:        bson.NewObjectId(),
	Username:  "祝绪丹",
	Password:  "666666",
	Age:       25,
	Interests: []string{"唱歌", "演戏", "跳舞"},
}

func main() {
	//----------------------------------- 新增 ------------------------------------
	//err := Insert(user1, user2) //插入数据，一次可以插入多个
	//if err != nil {
	//	log.Println(err)
	//}

	//----------------------------------- 删除 ------------------------------------
	//没有找到 err 为 not found
	//err := removeById() //更具 _id 删除一条
	//err := removeOneByField() //更具属性删除一条
	//_, err := removeAllByField() //更具属性删除多条
	//if err != nil {
	//	log.Println(err)
	//}

	//----------------------------------- 修改 ------------------------------------
	//update 没有找到 err为not found ；upset 没有找到会新增
	//err := updateById() //根据Id修改
	//_, err := upsetById() //根据Id修改,没有会新增
	//err := UpdateOneByField()    //修改满足条件的一条数据
	_, err := UpdateAllByField() //修改满足条件的所有条数据

	if err != nil {
		log.Println(err)
	}
	//----------------------------------- 查询 ------------------------------------
	//没有找到 err 为 not found
	err, data := FindAll() //查询全部
	//err, data := findById() //根据 _id 查询
	//err, data := findOneByField() //根据字段查询一条
	//err, data := findManyByField() //根据字段查询多条

	//----------------------------------------------------------------------------
	if err != nil {
		log.Println(err)
	}
	result, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(result))
	}
}

//---------------------------------------------------  插入数据 ---------------------------------------------------------
func Insert(data ...interface{}) error {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	for i := 0; i < len(data); i++ {
		err = db.Collection.Insert(data[i])
		if err != nil {
			return err
		}
	}
	return err
}

//---------------------------------------------------  删除数据 ---------------------------------------------------------
//根据 _id 删除数据
func removeById() error {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	err = db.Collection.Remove(bson.M{"_id": bson.ObjectIdHex("5dcec9ca43ea2b21ac51605d")})
	err = db.Collection.RemoveId(bson.ObjectIdHex("5dcec9ca43ea2b21ac51605e"))

	return err
}

//删除满足条件的一条
func removeOneByField() error {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	err = db.Collection.Remove(bson.M{"name": "小明"}) //删除满足条件的第一条
	//err = db.Collection.Remove(nil)  //修改条件为nil会删除第一条

	return err
}

//删除满足条件的多条
func removeAllByField() (*mgo.ChangeInfo, error) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	changeInfo, err := db.Collection.RemoveAll(bson.M{"name": "陈钰琪"}) //删除满足条件的多条
	//changeInfo,err := db.Collection.RemoveAll(nil)	//修改条件为nil会删除所有

	fmt.Println("匹配条数:", changeInfo.Matched, "删除条数:", changeInfo.Removed, "更新条数:", changeInfo.Updated, "增加条数:", changeInfo.UpsertedId)
	return changeInfo, err
}

//---------------------------------------------------  修改数据 (Upsert,Update用法相同)----------------------------------
//根据Id修改 没有 err为not found
func updateById() error {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	err = db.Collection.UpdateId(bson.ObjectIdHex("5dcec4afcfb0291c58d51084"), bson.M{"$set": bson.M{"name": "小刚", "pass": "pass"}})
	//err = db.Collection.Update(bson.M{"_id": bson.ObjectIdHex("5dcec4afcfb0291c58d51084")}, bson.M{"$set": bson.M{"name": "小红","pass": "111"}})

	return err
}

//通过Id修改 没有会将修改的字段新增为一条记录
func upsetById() (*mgo.ChangeInfo, error) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	changeInfo, err := db.Collection.UpsertId(bson.ObjectIdHex("5dcec4afcfb0291c58d51111"), bson.M{"$set": bson.M{"name": "小明", "pass": "pass"}})
	//changeInfo, err := db.Collection.Upsert(bson.M{"_id": bson.ObjectIdHex("5dcec4afcfb0291c58d52222")}, bson.M{"$set": bson.M{"name": "花花", "pass": "111"}})

	fmt.Println("匹配条数:", changeInfo.Matched, "删除条数:", changeInfo.Removed, "更新条数:", changeInfo.Updated, "增加条数:", changeInfo.UpsertedId)
	return changeInfo, err
}

//修改满足条件的一条数据
func UpdateOneByField() error {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	err = db.Collection.Update(bson.M{"name": "祝绪丹"}, bson.M{"$set": bson.M{"name": "小黑", "pass": "11234567"}}) //将 name=祝绪丹 的第一条数据进行修改name 和 pass
	//err = db.Collection.Update(nil, bson.M{"$set": bson.M{"name": "小黑", "pass": "11234567"}})  //修改条件为nil,会修改第一条

	return err
}

//修改满足条件的多条数据
func UpdateAllByField() (*mgo.ChangeInfo, error) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	changeInfo, err := db.Collection.UpdateAll(bson.M{"name": "祝绪丹"}, bson.M{"$set": bson.M{"name": "小慧", "pass": "11234567"}}) //将 name=祝绪丹 的所有数据进行修改name 和 pass
	//changeInfo, err := db.Collection.UpdateAll(nil, bson.M{"$set": bson.M{"name": "小黑", "pass": "11234567"}})  //修改条件为nil,会修改所有

	fmt.Println("匹配条数:", changeInfo.Matched, "删除条数:", changeInfo.Removed, "更新条数:", changeInfo.Updated, "增加条数:", changeInfo.UpsertedId)
	return changeInfo, err
}

//---------------------------------------------------  查询数据 ---------------------------------------------------------
//查询全部
func FindAll() (error, []User) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()
	var res []User

	err = db.Collection.Find(nil).All(&res)

	return err, res
}

//根据Id进行查询
func findById() (error, User) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()
	var res User

	err = db.Collection.FindId(bson.ObjectIdHex("5da08b1743ea2b15f473df29")).One(&res)

	return err, res
}

//根据普通字段查询(selector,filter可为nil)
/*
	Find(bson.M{"filedname":"value"}),查询条件
	Select({bson.M{"filedname":0})，表示忽略该字段则结果不返回此字段
	Select({bson.M{"filedname":1})，表示关注该字段则只返回关注字段
*/
func findOneByField() (error, User) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()
	var res User

	//err = db.Collection.Find(bson.M{"name": "小红"}).One(&res) //查询 name=小明 的完整数据
	err = db.Collection.Find(bson.M{"name": "小红"}).Select(bson.M{"_id": 0, "name": 1, "age": 1}).One(&res) //查询一条指定字段数据

	return err, res
}

//根据普通字段查询(selector,filter可为nil)
func findManyByField() (error, []User) {
	db := mongoUtils.DbConnection{DatebaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()
	var res []User

	//err = db.Collection.Find(nil).All(&res) //查询多条完整数据
	err = db.Collection.Find(bson.M{"name": "小明"}).Select(bson.M{"_id": 0, "name": 1, "age": 1}).All(&res) //查询多条指定字段数据 注意：多个Select().Select()后边的Select会覆盖前边

	return err, res
}
