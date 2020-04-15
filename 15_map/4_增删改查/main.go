package main

import (
	"fmt"
)

func main() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	//增加
	cities["no4"] = "济南"

	//修改(存在会覆盖)
	cities["no1"] = "qyz"
	fmt.Println(cities)

	//删除
	delete(cities, "no1")
	fmt.Println(cities)
	delete(cities, "no5") //当delete指定的key不存在时，删除不会操作，也不会报错
	fmt.Println(cities)

	//演示map的查找
	val, ok := cities["no1"]
	if ok {
		fmt.Printf("有key=no1,值为%v\n", val)
	} else {
		fmt.Printf("没有no1\n")
	}

	//如果希望一次性删除所有的key
	//1. 遍历所有的key,如何逐一删除 [遍历]
	//2. 直接make一个新的空间
	cities = make(map[string]string)
	fmt.Println(cities)

}
