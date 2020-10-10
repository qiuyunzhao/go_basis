package main

import (
	"fmt"
	"go_basis/34_数据库操作/1_Mongo/mongoUtils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//驱动官方文档： https://godoc.org/gopkg.in/mgo.v2

type User struct {
	Id        bson.ObjectId `bson:"_id"`
	Username  string        `bson:"name"`
	Password  string        `bson:"pass"`
	Age       int64         `bson:"age"`
	Interests []string      `bson:"interests"`
	Friends   []User        `bson:"friends"`
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
	Friends:   []User{user1},
}

var user3 = User{
	Id:        bson.NewObjectId(),
	Username:  "李沁",
	Password:  "88667",
	Age:       25,
	Interests: []string{"吃饭", "睡觉", "综艺"},
	Friends:   []User{user1, user2},
}

//---------------------------------------------------  插入数据 ---------------------------------------------------------
func Insert(data ...interface{}) error {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
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
func removeById(_id string) error {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	//err = db.Collection.Remove(bson.M{"_id": bson.ObjectIdHex(_id)})
	err = db.Collection.RemoveId(bson.ObjectIdHex(_id))

	return err
}

//删除满足条件的一条
func removeOneByField() error {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	err = db.Collection.Remove(bson.M{"name": "陈钰琪"}) //删除满足条件的第一条
	//err = db.Collection.Remove(nil)  //修改条件为nil会删除第一条

	return err
}

//删除满足条件的多条
func removeAllByField() (*mgo.ChangeInfo, error) {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
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
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	err = db.Collection.UpdateId(bson.ObjectIdHex("5e9662ae300118610c634486"), bson.M{"$set": bson.M{"name": "小刚", "pass": "pass"}})
	//err = db.Collection.Update(bson.M{"_id": bson.ObjectIdHex("5dcec4afcfb0291c58d51084")}, bson.M{"$set": bson.M{"name": "小红","pass": "111"}})

	return err
}

//通过Id修改 没有会将修改的字段新增为一条记录
func upsetById() (*mgo.ChangeInfo, error) {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	changeInfo, err := db.Collection.UpsertId(bson.ObjectIdHex("5dcec4afcfb0291c58d51111"), bson.M{"$set": bson.M{"name": "小明", "pass": "pass"}})
	//changeInfo, err := db.Collection.Upsert(bson.M{"_id": bson.ObjectIdHex("5dcec4afcfb0291c58d52222")}, bson.M{"$set": bson.M{"name": "花花", "pass": "111"}})

	fmt.Println("匹配条数:", changeInfo.Matched, "删除条数:", changeInfo.Removed, "更新条数:", changeInfo.Updated, "增加条数:", changeInfo.UpsertedId)
	return changeInfo, err
}

//修改满足条件的一条数据
func UpdateOneByField() error {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	err = db.Collection.Update(bson.M{"name": "祝绪丹"}, bson.M{"$set": bson.M{"name": "小黑", "pass": "11234567"}}) //将 name=祝绪丹 的第一条数据进行修改name 和 pass
	//err = db.Collection.Update(nil, bson.M{"$set": bson.M{"name": "小黑", "pass": "11234567"}})  //修改条件为nil,会修改第一条

	return err
}

//修改满足条件的多条数据
func UpdateAllByField() (*mgo.ChangeInfo, error) {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	changeInfo, err := db.Collection.UpdateAll(bson.M{"name": "祝绪丹"}, bson.M{"$set": bson.M{"name": "小慧", "pass": "11234567"}}) //将 name=祝绪丹 的所有数据进行修改name 和 pass
	//changeInfo, err := db.Collection.UpdateAll(nil, bson.M{"$set": bson.M{"name": "小黑", "pass": "11234567"}})  //修改条件为nil,会修改所有

	fmt.Println("匹配条数:", changeInfo.Matched, "删除条数:", changeInfo.Removed, "更新条数:", changeInfo.Updated, "增加条数:", changeInfo.UpsertedId)
	return changeInfo, err
}

// 嵌套更改 https://blog.csdn.net/liuchangqing123/article/details/48106493
func nestedUpdate() error {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()

	// 修改 friends[0] 的 pass
	//_, err = db.Collection.UpdateAll(bson.M{"friends.friends.name": "陈钰琪", "friends.friends.age": 24}, bson.M{"$set": bson.M{"friends.0.pass": "667788"}})
	//if err != nil {
	//	return err
	//}

	// $ 通配符 匹配数组的多个 不能用在最后一级
	_, err = db.Collection.UpdateAll(bson.M{"friends.friends.name": "陈钰琪", "friends.friends.age": 24}, bson.M{"$set": bson.M{"friends.$.friends.0.pass": "02"}})
	if err != nil {
		return err
	}

	return nil
}

//---------------------------------------------------  查询数据 查询不到返回 err, err.Error() 为"not found" ---------------------------------------------------------
//查询全部
func FindAll() (error, []User) {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()
	var res []User

	err = db.Collection.Find(nil).All(&res)

	return err, res
}

//根据Id进行查询
func findById() (error, User) {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()
	var res User

	err = db.Collection.FindId(bson.ObjectIdHex("5e9664c830011837a4963085")).One(&res)

	return err, res
}

//根据普通字段查询(selector,filter可为nil)
/*
	Find(bson.M{"filedname":"value"}),查询条件
	Select({bson.M{"filedname":0})，表示忽略该字段则结果不返回此字段
	Select({bson.M{"filedname":1})，表示关注该字段则只返回关注字段
*/
func findOneByField() (error, User) {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()
	var res User

	//err = db.Collection.Find(bson.M{"name": "陈钰琪"}).One(&res) //查询 name=陈钰琪 的完整数据
	err = db.Collection.Find(bson.M{"name": "陈钰琪"}).Select(bson.M{"_id": 0, "name": 1, "age": 1}).One(&res) //查询一条指定字段数据

	return err, res
}

//根据普通字段查询(selector,filter可为nil)
func findManyByField() (error, []User) {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()
	var res []User

	//err = db.Collection.Find(nil).All(&res) //查询多条完整数据
	err = db.Collection.Find(bson.M{"name": "陈钰琪"}).Select(bson.M{"_id": 0, "name": 1, "age": 1}).All(&res) //查询多条指定字段数据 注意：多个Select().Select()后边的Select会覆盖前边

	return err, res
}

// 嵌套查询
func nestedQuery() (error, []User) {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()
	var users []User

	err = db.Collection.Find(bson.M{"friends.friends.name": "陈钰琪", "friends.friends.age": 24}).Select(bson.M{"friends.friends": 1}).All(&users) //查询一条指定字段数据

	return err, users
}

// 查询条件或运算
// https://blog.csdn.net/LightUpHeaven/article/details/82663146
// https://blog.csdn.net/tianwenxue/article/details/106316255
func MulticonditionalQuery() (error, []User) {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()
	var users []User

	err = db.Collection.Find(bson.D{{"$or", []interface{}{bson.D{{"name", "祝绪丹"}}, bson.D{{"name", "李沁"}}}}}).All(&users)

	return err, users
}

//根据字段模糊查询
//https://blog.csdn.net/Nick_666/article/details/81286239?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.nonecase
//option的值包含：
//  i(不区分大小写)，
//  m(当使用^与$符号模糊匹配时，作用于屏蔽中间的换行符) ,
//  x(忽略注释，以#开头 /n结尾)，
//  s(允许所有字符包括换行符参与模糊匹配)
func LikeFindByField() (error, []User) {
	db := mongoUtils.DbConnection{DatabaseName: databaseName, CollectionName: collectionName}
	err := db.ConnDB()
	defer db.CloseDB()
	var res []User

	//err = db.Collection.Find(bson.M{"age": 25, "name": "陈钰琪"}).All(&res) // 条件查询
	err = db.Collection.Find(bson.M{"age": 25, "name": bson.M{"$regex": "陈", "$options": "$i"}}).All(&res) // 与运算的条件查询
	return err, res
}
