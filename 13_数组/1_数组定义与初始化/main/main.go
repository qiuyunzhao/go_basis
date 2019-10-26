package main

import (
	"fmt"
)

//Go的数组属值类型， 在默认情况下是值传递
func main() {

	var intArr [3]int   //int占8个字节
	fmt.Println(intArr) //当我们定义完数组后，其实数组的各个元素有默认值 0

	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)

	fmt.Printf("intArr的地址=%p intArr[0] 地址%p intArr[1] 地址%p intArr[2] 地址%p\n",
		&intArr, &intArr[0], &intArr[1], &intArr[2]) //int占8字节，数组地址是连续的

	//四种初始化数组的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=", numArr01)

	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr02=", numArr02)

	var numArr03 = [...]int{8, 9, 10} //这里的 [...] 是规定的写法
	fmt.Println("numArr03=", numArr03)

	var numArr04 = [...]int{1: 800, 0: 900, 2: 999}
	fmt.Println("numArr04=", numArr04)

	//类型推导
	strArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Println("strArr05=", strArr05)

	//GO中 数组是值类型
	arr := [3]int{11, 22, 33}
	fmt.Printf("arr的地址=%p \n", &arr)
	test01(arr)
	fmt.Println("基础用法 arr=", arr)
	test02(&arr)
	fmt.Println("基础用法 arr=", arr)
}

//[3]int、[4]int、[]int 是三种不同的数据类型，要注意长度数组类型的一部分
func test01(arr [3]int) {
	fmt.Printf("arr的地址=%p \n", &arr)
	(arr)[0] = 88
}

func test02(arr *[3]int) {
	fmt.Printf("arr指针的地址=%p \n", &arr)
	(*arr)[0] = 88
}
