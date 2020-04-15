package main

import "fmt"

//切片是引用类型的，底层是通过指向一个数组的引用来实现的
func main() {
	//声明并定义一个数组
	var intArr = [...]int{1, 22, 33, 66, 99}

	//方式1
	//intArr[1:3] 表示 slice 引用到intArr这个数组，引用intArr数组的起始下标为 1 , 最后的下标为3(但是不包含3)
	slice1 := intArr[1:3]
	fmt.Println("slice 的元素是 =", slice1)       //[22 33]
	fmt.Println("slice 的元素个数 =", len(slice1)) // 2
	fmt.Println("slice 的容量 =", cap(slice1))   // 4(这只是初始化的容量，切片的容量是可以动态变化)

	//切片是引用类型
	fmt.Printf("intArr[1]的地址=%p\n", &intArr[1]) //0xc00008e038
	fmt.Printf("slice[0]的地址=%p \n", &slice1[0]) //0xc00008e038
	test(slice1)
	fmt.Println("intArr=", intArr)      //[1 100 33 66 99]
	fmt.Println("slice 的元素是 =", slice1) //[100 33]

	//方式2：:使用 make
	var slice2 = make([]float64, 5, 10)
	slice2[1] = 10
	slice2[3] = 20
	fmt.Println(slice2)

	//方式3：定义一个切片，直接就指定具体数组
	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strSlice=", strSlice)

}

func test(slice []int) {
	slice[0] = 100 //这里修改slice[0],会改变实参
}
