package main

import (
	"fmt"
)

func main() {
	var arr = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4] // [20, 30, 40]

	fmt.Println("-------------------------------- 1.使用常规的for循环遍历切片 -------------------------------------")

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v \n", i, slice[i])
	}

	fmt.Println("------------------------------ 2.使用for--range 方式遍历切片 -------------------------------------")

	for i, v := range slice {
		fmt.Printf("i=%v v=%v \n", i, v)
	}

	fmt.Println("------------------------------------ 3.将切片再进行切片 ------------------------------------------")

	slice2 := slice[1:2] //  slice [20, 30, 40]    slice2 [30]
	slice2[0] = 100      // 因为arr , slice 和slice2 指向的数据空间是同一个，因此slice2[0]=100，其它的都变化

	fmt.Println("slice2=", slice2) //[100]
	fmt.Println("slice=", slice)   //[20 100 40]
	fmt.Println("arr=", arr)       //[10 20 100 40 50]

	fmt.Println("---------------------------- 4.append内置函数，对切片进行动态追加 ---------------------------------")

	var slice3 = []int{100, 200, 300}
	//通过append直接给slice3追加具体的元素
	slice3 = append(slice3, 400, 500, 600)
	fmt.Println("slice3", slice3) //[100 200 300 400 500 600]

	//通过append将切片slice3追加给slice3
	slice3 = append(slice3, slice3...) //[100 200 300 400 500 600 100 200 300 400 500 600]
	fmt.Println("slice3", slice3)

	fmt.Println("------------------------------- 5.opy内置函数，对切片进行复制 ------------------------------------")
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	var slice6 = make([]int, 2)
	copy(slice5, slice4)
	copy(slice6, slice4)

	fmt.Println("slice4=", slice4) // [1 2 3 4 5]
	fmt.Println("slice5=", slice5) // [1 2 3 4 5 0 0 0 0 0]
	fmt.Println("slice6=", slice6) // [1 2]  会截取
}
