package main

import "fmt"

//获取控制台输入的 【姓名,年龄,薪水,是否是男生】
func main() {
	var name string
	var age byte
	var salary float32
	var isMan bool

	//方式1
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&salary)
	fmt.Println("请输入是否是男生")
	fmt.Scanln(&isMan)

	fmt.Println(" 名字：", name, "\n 年龄：", age, "\n 薪水：", salary, "\n 是否为男生：", isMan)

	//方式2
	fmt.Println("请输入 [姓名,年龄,薪水,是否是男生] ")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isMan)
	fmt.Println(" 名字：", name, "\n 年龄：", age, "\n 薪水：", salary, "\n 是否为男生：", isMan)
}
