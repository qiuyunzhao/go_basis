package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//test01() //写的方式打开文件，清空原文件内容再写入
	//test02() //写的方式打开文件，追加的方式写入
	//test03() //读写的方式打开文件，追加的方式写入
	test04() //从一个人文件读取内容写入另一个文件
}

//写的方式打开文件，清空原文件内容再写入
func test01() {
	filePath := "go_code/18_文件操作/sources/abc.txt"

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666) //清空原文件内容再写入
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	str := "清空原文件内容再写入\r\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}

//写的方式打开文件，追加的方式写入
func test02() {
	filePath := "go_code/18_文件操作/sources/abc.txt"

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 6666) //追加的方式写入
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	str := "追加的方式写入\r\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}

//读写的方式打开文件，追加的方式写入
func test03() {
	filePath := "go_code/18_文件操作/sources/abc.txt"

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666) //读写的方式打开文件，追加的方式写入
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	str := "hello,北京!\r\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}

//从一个人文件读取内容写入另一个文件
func test04() {
	fromFilePath := "go_code/18_文件操作/sources/a.txt"
	toFilePath := "go_code/18_文件操作/sources/abc.txt"

	data, err := ioutil.ReadFile(fromFilePath)
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
		return
	}

	err = ioutil.WriteFile(toFilePath, data, 0666)
	if err != nil {
		fmt.Printf("write file error=%v\n", err)
	}
}
