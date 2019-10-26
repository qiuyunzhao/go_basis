package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//自己编写一个函数，接收两个文件路径 srcFileName dstFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	//通过srcfile ,获取到 Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName(不存在就新建)
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer dstFile.Close()
	//通过dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)

	//拷贝
	return io.Copy(writer, reader)
}

func main() {
	srcFile := "18_文件操作/sources/赵敏.jpg"
	dstFile := "18_文件操作/sources/赵敏copy.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Printf("拷贝完成\n")
	} else {
		fmt.Printf("拷贝错误 err=%v\n", err)
	}

}
