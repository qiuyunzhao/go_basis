package main

import (
	"fmt"
)

func main() {
	heroes := [...]string{"宋江", "吴用", "卢俊义"}

	//使用常规的方式遍历
	for i := 0; i < len(heroes); i++ {
		fmt.Printf("元素%d的值：%v \n", i, heroes[i])
	}

	//演示for-range遍历数组
	for index, value := range heroes {
		fmt.Printf("heroes[%d]=%v \n", index, value)
	}

	for _, value := range heroes {
		fmt.Printf("元素的值=%v\n", value)
	}
}
