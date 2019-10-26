package main

import (
	"fmt"

	//完整包名 "D:\GoProjects\src\go_code\3_跨包引用\model"
	// GoPath为：D:\GoProjects可省略
	// 默认 \src 目录，可省略
	"go_code/3_跨包引用/model"

	//为包起别名，本报内引用包之前的包名不在生效
	Myutils "go_code/3_跨包引用/utils"
)

func main() {
	//fmt.Println(model.hero)
	fmt.Println(model.Hero)

	//model.printHero()
	model.PrintHero()

	Myutils.Printutils()

}
