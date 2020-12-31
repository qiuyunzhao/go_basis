package model

import "fmt"

var hero string = "钢铁侠"  //小写不能跨包引用
var Hero string = "美国队长" //大写能跨包引用

func printHero() {
	fmt.Println(hero)
}

func PrintHero() {
	fmt.Println(hero)
}
