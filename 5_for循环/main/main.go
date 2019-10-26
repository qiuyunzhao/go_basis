package main

import "fmt"

//go语言没有while  和  do while 语句，可以用for循环代替
func main() {

	//替换while
	var i int = 1
	for {
		if i > 5 {
			break
		}
		fmt.Println("hello world", i)
		i++
	}

	//替换do while
	var j int = 1
	for {
		fmt.Println("hi world", j)
		j++
		if j > 5 {
			break
		}
	}

	var str string = "hi!北京"
	//字符串遍历方式1-传统方式(按字节遍历，中文会乱码)
	for i = 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i]) //适用下标
	}
	fmt.Println()

	//用切片解决乱码问题
	str2 := []rune(str)
	for i = 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i]) //适用下标
	}
	fmt.Println()

	//字符串遍历方式2 - for range()
	for index, val := range str {
		fmt.Printf("index=%d,val=%c \n", index, val)
	}

}
