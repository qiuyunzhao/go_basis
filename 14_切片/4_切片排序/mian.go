package main

import (
	"fmt"
	"sort"
)

type people struct {
	Name    string
	Age     int
	Address string
}

//https://blog.csdn.net/yzf279533105/article/details/89459460
func main() {
	//普通类型切片排序
	testIntSort()
	testStrings()
	testFloat()
	testIntSearch()

	//复杂类型切片排序
	s := make([]people, 6)
	s[0] = people{Name: "张无忌", Age: 19, Address: "光明顶"}
	s[1] = people{Name: "孙悟空", Age: 500, Address: "花果山"}
	s[2] = people{Name: "张三丰", Age: 105, Address: "武当山"}
	s[3] = people{Name: "刘老根", Age: 29, Address: "东北"}
	s[4] = people{Name: "刘备", Age: 59, Address: "荆州"}
	s[5] = people{Name: "张无忌", Age: 59, Address: "嵩山"}
	fmt.Println("初始化结果--------------", s)

	//注意下边是引用传递 而不是值传递
	testSlice(s)
	testSliceStable(s)
	testSliceIsSorted(s)
	testOrderNameFirstAndAgeNext(s)
}

// ------------------------------------------- 普通类型切片排序 ----------------------------------------------------
func testIntSort() {
	var a = [...]int{1, 8, 38, 2, 348, 484} //定义数组
	sort.Ints(a[:])                         //数组是值类型,不能直接排序，必须转为切片
	fmt.Println(a)
}

func testStrings() {
	var a = [...]string{"abc", "efg", "b", "A", "eeee"} //定义数组
	sort.Strings(a[:])                                  //按照字母顺序排序,从小到大
	fmt.Println(a)
}

func testFloat() {
	var a = [...]float64{2.3, 0.8, 28.2, 392342.2, 0, 6}
	sort.Float64s(a[:])
	fmt.Println(a)
}

// ------------------------------------------- 排序后查找位置 ----------------------------------------------------
func testIntSearch() {
	var a = [...]int{1, 8, 38, 2, 348, 484}
	sort.Ints(a[:])
	index := sort.SearchInts(a[:], 348) //SearchInts默认排序后的位置，一定要排序后在查找
	fmt.Println(index)
}

// ------------------------------------------- 结构体类型切片排序 ----------------------------------------------------
//  https://blog.csdn.net/weixin_42900065/article/details/97758432

//  go语言的slice()不仅仅可以对int类型的数组进行排序，还可以对struct类型的数组进行排序,排序函数如下
//		1. Slice() 不稳定排序     -- 按某一规则排序，相等的元素老的切片中的顺序可能变化(一般不会有影响，除非特殊需求)
//		2. SliceStable() 稳定排序 -- 按某一规则排序，相等的元素老的切片中的顺序不会变化
//		3. SliceIsSorted() 判断是否已排序

//1. Slice() 不稳定排序
func testSlice(s []people) {
	// 从小到大排序
	sort.Slice(s, func(i, j int) bool {
		if s[i].Age < s[j].Age {
			return true
		}
		return false
	})
	fmt.Println("从小到大不稳定排序结果---", s)

	// 从大到小排序
	sort.Slice(s, func(i, j int) bool {
		if s[i].Age > s[j].Age {
			return true
		}
		return false
	})
	fmt.Println("从大到小不稳定排序结果---", s)
}

//2. SliceStable() 稳定排序
func testSliceStable(s []people) {
	// 从小到大排序
	sort.SliceStable(s, func(i, j int) bool {
		if s[i].Age < s[j].Age {
			return true
		}
		return false
	})
	fmt.Println("从小到大稳定排序结果-----", s)

	// 从大到小排序
	sort.SliceStable(s, func(i, j int) bool {
		if s[i].Age > s[j].Age {
			return true
		}
		return false
	})
	fmt.Println("从大到小稳定排序结果-----", s)
}

//3. SliceIsSorted() 判断是否已排序
func testSliceIsSorted(s []people) {
	// 是否从小到大排序
	bLess := sort.SliceIsSorted(s, func(i, j int) bool {
		if s[i].Age < s[j].Age {
			return true
		}
		return false
	})
	fmt.Printf("数组s是否从小到大排序:%v\n", bLess)

	// 是否从大到小排序
	bLarger := sort.SliceIsSorted(s, func(i, j int) bool {
		if s[i].Age > s[j].Age {
			return true
		}
		return false
	})
	fmt.Printf("数组s是否从大到小排序:%v\n", bLarger)
}

//首先按照年龄排序然后按照名字排序
func testOrderNameFirstAndAgeNext(s []people) {
	// 从小到大排序
	sort.Slice(s, func(i, j int) bool {
		if s[i].Name < s[j].Name {
			return true
		}
		if s[i].Name == s[j].Name {
			if s[i].Age < s[j].Age {
				return true
			}
		}
		return false
	})

	fmt.Println("先按名字后按年龄从小到大排序结果--", s)
}
