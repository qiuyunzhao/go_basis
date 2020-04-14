package main

import (
	"fmt"
	"go_basis/16_结构体（似java的类）/6_封装（工厂模式）/model"
)

func main() {
	//创建Student实例
	// var stu = model.Student{
	// 	Name :"tom",
	// 	Score : 78.9,
	// }

	//定student结构体是首字母小写，我们可以通过工厂模式来解决
	var stu = model.NewStudent("tom~", 98.8)

	fmt.Println(stu)                                          //&{tom~ 98.8}
	fmt.Println("name=", stu.Name, " score=", stu.GetScore()) //name= tom~  score= 98.8
}
