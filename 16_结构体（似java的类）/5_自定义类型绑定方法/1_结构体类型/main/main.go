package main

//自定义类型都可以绑定方法，不限于结构体  type Integer int  的自定义类型 Integer 也可以绑定方法
import (
	"fmt"
)

type Person struct {
	Name string
}

//为自定义类型绑定方法，绑定的方法只能通过该类型的变量来调用（在方法的栈区会值拷贝一份 p Person 相当于值类型传递）
func (p Person) test03() {
	p.Name = "jack"
	fmt.Println("test03() =", p.Name) // jack
}

//通常采用指针传递，提高效率(在方法的栈区会值传入 p Person 的地址 相当于指针类型传递)
func (p *Person) test04() {
	p.Name = "mary"
	fmt.Println("test03() =", p.Name) // mary
}

//给Person结构体添加getSum方法,可以计算两个数的和，并返回结果
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {

	p := Person{"tom"}

	//结构体方法的值传递
	p.test03()
	fmt.Println("基础用法() p.name=", p.Name) // tom

	(&p).test03()                         // 从形式上是传入地址，但是本质仍然是值拷贝
	fmt.Println("基础用法() p.name=", p.Name) // tom

	//结构体方法的指针传递
	p.Name = "tom"
	(&p).test04()
	fmt.Println("基础用法() p.name=", p.Name) // mary

	p.Name = "tom"
	p.test04()                            // 等价(&p).test04,从形式上是传入值类型，但是本质仍然是地址拷贝（go为了使用方便，编译器会自己加上）
	fmt.Println("基础用法() p.name=", p.Name) // mary

	sum := p.getSum(10, 20)
	fmt.Println("sum=", sum)

}
