package main

import (
	"fmt"
	"os"
)

//概念说明: file 的叫法
//1. file 叫 file对象
//2. file 叫 file指针
//3. file 叫 file 文件句柄

func main() {
	//打开文件
	file, err := os.Open("go_code/18_文件操作/sources/a.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//输出下文件，看看文件是什么, 看出file 就是一个指针 *File
	fmt.Printf("file= %v", file) //file=&{0xc000090780}

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}

}
