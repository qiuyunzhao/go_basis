package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

//声明一个Hero结构体切片类型
type HeroSlice []Hero

//切片实现Interface 接口的三个方法
//1.Len() ：返回切片的大小
func (hs HeroSlice) Len() int {
	return len(hs)
}

//2.Less(i, j int) :决定使用什么规则进行排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age // 按Hero的Age从小到大排序
	//return hs[i].Name > hs[j].Name  // 按Hero的Name从大到小排序
}

//3.Swap(i, j int) :Less(i, j int)返回true时进行交换
func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	//---------------------------------------------- 基本类型切片排序 --------------------------------------------------
	var intSlice = []int{0, -1, 10, 7, 90}
	//使用系统提供的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//---------------------------------------------- 结构体切片排序 ---------------------------------------------------
	var heroes HeroSlice
	//初始化切片
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}

	//排序前的顺序
	fmt.Println("-----------排序前------------")
	for _, v := range heroes {
		fmt.Println(v)
	}

	//结构体实现了排序Interface接口,可以调用sort.Sort进行结构体排序
	sort.Sort(heroes)

	fmt.Println("-----------排序后------------")
	//看看排序后的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}

}
