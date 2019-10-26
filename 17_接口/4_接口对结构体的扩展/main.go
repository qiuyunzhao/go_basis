package main

//接口可以在不破坏结构体的情况下对继承结构体的子类进行扩展

import (
	"fmt"
)

//结构体父类
type Animal struct {
	Name string
}

func (this *Animal) Running() {
	fmt.Println(this.Name, "天生会跑")
}

//接口
type MonkeyAble interface {
	Climbing()
}

type FishAble interface {
	Swimming()
}

//子结构体
type People struct {
	Animal //继承
}

//让People实现MonkeyAble
func (this *People) Climbing() {
	fmt.Println(this.Name, "通过学习，会爬树...")
}

//让People实现FishAble
func (this *People) Swimming() {
	fmt.Println(this.Name, "通过学习，会游泳..")
}

func main() {

	person := People{}
	person.Name = "小明"
	person.Running()
	person.Climbing()
	person.Swimming()

}
