package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `bson:"_id"`
	Username  string        `bson:"name"`
	Password  string        `bson:"pass"`
	Age       int64         `bson:"age"`
	Interests []string      `bson:"interests"`
}

type Student struct {
	Id    bson.ObjectId `bson:"_id"`
	Name  string        `bson:"name"`
	Phone string        `bson:"phone"`
	Email string        `bson:"email"`
	Sex   string        `bson:"sex"`
}

const (
	url      = "127.0.0.1:27017"
	database = "userinfo"
)

func main() {

	//连接数据库
	database, err := ConnnectDB()
	if err == nil {
		fmt.Println("连接数据库成功")
	}

	//data := User{
	//	Id:        bson.NewObjectId(),
	//	Username:  "小明",
	//	Password:  "666666",
	//	Age:       24,
	//	Interests: []string{"象棋", "游泳", "跑步"},
	//}
	//stu := Student{
	//	Id:    bson.NewObjectId(),
	//	Name:  "陈钰琪",
	//	Phone: "17615110273",
	//	Email: "1102401880@qq.com",
	//	Sex:   "女",
	//}
	//err = add(database, "user", data, stu)
	//if err == nil {
	//	fmt.Println("插入成功")
	//}

	//err = removeById(database, "user", "5d9fea393001182cc8c7b0d9")
	//err = removeOneByField(database, "user", bson.M{"name": "陈钰琪"})
	//changeInfo, err := removeAllByField(database, "user", bson.M{"name": "陈钰琪"})
	//fmt.Println("匹配条数:", changeInfo.Matched, "删除条数:", changeInfo.Removed, "更新条数:", changeInfo.Updated, "增加条数:", changeInfo.UpsertedId)
	//if err == nil {
	//	fmt.Println("删除成功")
	//}

	//err = updateById(database, "user", "5d9feb3730011825c0f9c635",bson.M{"$set": bson.M{"name": "修改后的name","pass": "修改后的pass"}})
	//err = UpdateOneByField(database, "user", bson.M{"name": "小明"}, bson.M{"$set": bson.M{"name": "祝绪丹","pass": "8888888"}})
	//changeInfo, err := UpdateAllByField(database, "user", bson.M{"name": "小明"}, bson.M{"$set": bson.M{"name": "祝绪丹","pass": "8888888"}})
	//fmt.Println("匹配条数:", changeInfo.Matched, "删除条数:", changeInfo.Removed, "更新条数:", changeInfo.Updated, "增加条数:", changeInfo.UpsertedId)
	//if err == nil {
	//	fmt.Println("更新成功")
	//}

	//user := new(User)
	var users []User
	//err = findById(database, "user", "5d9ff5f130011858f0332b61",&user)
	//err = findByField(database, "user", bson.M{"name": "小明"}, bson.M{"_id": 0, "name": 1, "age": 1}, &user)
	err = findByField(database, "user", bson.M{"name": "小明"}, bson.M{"_id": 0, "name": 1, "age": 1}, &users) //切片只能是 &users传参
	//err = findAll(database, "user", &users)

	if err == nil {
		fmt.Println("查询成功")
	}

	//fmt.Println(user)
	for index, value := range users {
		fmt.Println(index, ":  ", value)
	}
}

//--------------------------------------- 连接数据库获取要插入的集合（表） ----------------------------------------------
func ConnnectDB() (*mgo.Database, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	session.SetMode(mgo.Monotonic, true)
	//切换到数据库
	database := session.DB(database)
	return database, err
}

//---------------------------------------------------  新增数据 ---------------------------------------------------------
func add(database *mgo.Database, collection string, data ...interface{}) error {
	c := database.C(collection)
	err := c.Insert(&data[0], &data[1])
	return err
}

//---------------------------------------------------  删除数据 ---------------------------------------------------------
func removeById(database *mgo.Database, collection string, id string) error {
	c := database.C(collection)
	//err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	err := c.RemoveId(bson.ObjectIdHex(id))
	return err
}

func removeOneByField(database *mgo.Database, collection string, selector interface{}) error {
	c := database.C(collection)
	//删除满足条件的第一条
	err := c.Remove(selector)
	//修改条件为nil会删除第一条
	//err := c.Remove(nil)
	return err
}

func removeAllByField(database *mgo.Database, collection string, selector interface{}) (*mgo.ChangeInfo, error) {
	c := database.C(collection)
	//删除满足条件的多条
	changeInfo, err := c.RemoveAll(selector)
	//修改条件为nil会删除所有
	//changeInfo,err := c.RemoveAll(nil)
	return changeInfo, err
}

//---------------------------------------------------  修改数据 ---------------------------------------------------------
//通过Id修改
func updateById(database *mgo.Database, collection string, id string, update interface{}) error {
	c := database.C(collection)
	err := c.UpdateId(bson.ObjectIdHex(id), update)
	//err := c.Update(bson.M{"_id": objectId}, update)
	return err
}

//根据普通字段修改
func UpdateOneByField(database *mgo.Database, collection string, selector interface{}, update interface{}) error {
	c := database.C(collection)
	//修改满足条件的第一条
	err := c.Update(selector, update)
	//修改条件为nil会修改第一条
	//err := c.Update(nil, update)
	return err
}

//修改满足条件的多条数据
func UpdateAllByField(database *mgo.Database, collection string, selector interface{}, update interface{}) (*mgo.ChangeInfo, error) {
	c := database.C(collection)
	//修改满足条件的多条
	changeInfo, err := c.UpdateAll(selector, update)
	////修改条件为nil会修改所有
	//changeInfo, err := c.UpdateAll(nil, update)
	return changeInfo, err
}

//---------------------------------------------------  查询数据 ---------------------------------------------------------
//根据ObjectId进行查询
func findById(database *mgo.Database, collection string, id string, result interface{}) error {
	c := database.C(collection)
	err := c.FindId(bson.ObjectIdHex(id)).One(result)
	return err
}

//根据普通字段查询(selector,filter可为nil)
func findByField(database *mgo.Database, collection string, selector interface{}, filter interface{}, result interface{}) error {
	c := database.C(collection)

	/*
		Find(bson.M{"filedname":"value"}),查询条件
		Select({bson.M{"filedname":0})，表示忽略该字段则结果不返回此字段
		Select({bson.M{"filedname":1})，表示关注该字段则只返回关注字段
	*/

	//err := c.Find(selector).One(result) //查询一条完整数据
	//err := c.Find(selector).Select(filter).One(result) //查询一条指定字段数据

	//err := c.Find(selector).All(result) //查询多条完整数据
	err := c.Find(selector).Select(filter).All(result) //查询多条指定字段数据 注意：多个Select().Select()后边的Select会覆盖前边

	return err
}

//查询所有
func findAll(database *mgo.Database, collection string, result interface{}) error {
	c := database.C(collection)
	err := c.Find(nil).All(result)
	return err
}
