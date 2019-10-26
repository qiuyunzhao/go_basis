package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	//方式1
	var p1 Person
	p1.Name = "tom"
	p1.Age = 18
	fmt.Println(p1)

	//方式2
	p2 := Person{"mary", 20}
	p2 = Person{
		Name: "mary~",
		Age:  20,
	}
	fmt.Println(p2)

	//方式3-& 返回指针
	var p3 *Person = new(Person)
	//因为p3是一个指针，因此标准的给字段赋值方式  (*p3).Name = "smith" 也可以这样写 p3.Name = "smith"
	//原因: go的设计者为了程序员使用方便，底层会对 p3.Name = "smith" 进行处理，给 p3 加上 取值运算 (*p3).Name = "smith"
	(*p3).Name = "smith"
	p3.Name = "smith~"
	(*p3).Age = 30
	p3.Age = 100
	fmt.Println(*p3)

	//方式4-{} 返回结构体的指针类型
	//var person *Person = &Person{}
	var p4 *Person = &Person{"mary", 60}

	//p4 是一个指针，因此标准的访问字段的方法 (*person).Name = "scott"
	// go的设计者为了程序员使用方便，也可以 person.Name = "scott"，底层会对 person.Name = "scott" 进行处理， 会加上 (*person)
	(*p4).Name = "scott"
	p4.Name = "scott~"
	(*p4).Age = 88
	p4.Age = 10
	fmt.Println(*p4)

}
