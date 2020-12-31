package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

//参考：https://www.cnblogs.com/wdliu/p/9330278.html
func main() {
	//链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭..（一定要关闭）

	//操作字符串数据
	testStringOperate(conn)

	//操作字符串数据
	//testHashOperate(conn)

}

//操作字符串数据
func testStringOperate(conn redis.Conn) {
	//1. 通过go 向redis写入数据 string [key-val]
	result, err := conn.Do("Set", "name", "tomjerry猫猫") // key:name value:tomjerry猫猫    key相同会被覆盖
	if err != nil {
		log.Println("set  err=", err)
		return
	}
	fmt.Println("插入结果： ", result)

	_, err = conn.Do("expire", "name", 10) //10秒过期
	if err != nil {
		fmt.Println("set expire error: ", err)
		return
	}

	//2. 通过go 从redis读取数据 string [key-val]
	//因为conn.Do()返回 结果是 interface{} 且name对应的值是string ,因此我们需要转换  直接用 nameString := r.(string)转换不可以
	result, err = redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("set  err=", err)
		return
	}
	fmt.Println("查询结果： ", result)
}

//操作Hash数据
func testHashOperate(conn redis.Conn) {
	//1. 通过go向redis写入数据 string [key-val]
	result, err := conn.Do("HMSet", "user02", "name", "john", "age", 19, "address", "济南")
	if err != nil {
		log.Println("HMSet  err=", err)
		return
	}
	fmt.Println("插入结果： ", result)

	//2. 通过go从redis读取数据
	result1, err := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("hget  err=", err)
		return
	}
	fmt.Println("查询结果： ", result1)
	for i, v := range result1 {
		fmt.Printf("r[%d]=%s\n", i, v)
	}
}
