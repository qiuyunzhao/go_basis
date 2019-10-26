package main

import (
	"fmt"
)

func main() {
	//第一种方式
	var a map[string]string
	//在使用map前，需要先make , make的作用就是给map分配数据空间
	a = make(map[string]string, 10) //分配数据空间10
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松"
	a["no3"] = "吴用"
	fmt.Println(a)

	//第二种方式
	cities := make(map[string]string) //系统自己分配数据空间
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	//第三种方式
	heroes := map[string]string{
		"hero1": "宋江",
		"hero2": "卢俊义",
		"hero3": "吴用",
	}
	heroes["hero4"] = "林冲"
	heroes["hero1"] = "qyz" //相同的键，对应值会被覆盖
	fmt.Println("heroes=", heroes)

}
