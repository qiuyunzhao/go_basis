package mongoUtils

import (
	"gopkg.in/mgo.v2"
	"log"
)

const (
	MONGODB_URL = "127.0.0.1:27017"
	//MONGODB_URL="mongodb://admin:langchao!Test6530@10.24.20.71:27017" //绑定端口到28081
)

var (
	// 全局session连接池，默认最大连接数4096；
	// 利用session.Clone()或session.Copy()从连接池获取session连接；
	// 利用session.Close()将使用完的session连接,放回连接池；
	// 注意：使用完后一定要关闭否则达到最大连接数后会阻塞；
	session *mgo.Session
)

type DbConnection struct {
	DatabaseName   string
	CollectionName string
	Session        *mgo.Session
	Database       *mgo.Database
	Collection     *mgo.Collection
}

// 获取session，已经连接成功就复制一份，否则创建新的
func (db *DbConnection) getSession() error {
	if session == nil {
		var err error
		session, err = mgo.Dial(MONGODB_URL)
		if err != nil {
			log.Println("获取mongo数据库session连接失败!:", err.Error())
			session = nil
			return err
		}
		log.Println("新建session---------------------------")
	}
	db.Session = session.Clone()
	return nil
}

// 连接数据库和数据表
func (db *DbConnection) ConnDB() error {
	err := db.getSession()
	if err != nil {
		return err
	}
	db.Session.SetMode(mgo.Eventual, true) //不缓存连接模式
	db.Database = db.Session.DB(db.DatabaseName)
	db.Collection = db.Database.C(db.CollectionName)
	return nil
}

// 关闭数据连接的session
func (db *DbConnection) CloseDB() {
	if db.Session != nil {
		db.Session.Close()
	}
}
