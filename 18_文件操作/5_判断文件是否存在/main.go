package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file := "18_文件操作/sources/赵敏.jpg"

	IsExist := IsExist(file)
	fmt.Println(IsExist)

	if !IsExist {
		// 递归创建文件夹
		if err := os.MkdirAll("18_文件操作/sources", os.ModePerm); err != nil {
			log.Println("os.MkdirAll错误:", err)
		}
	}

	fmt.Printf("%s is file: %v\n", file, IsFile(file))
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsFile checks whether the path is a file,
// it returns false when it's a directory or does not exist.
func IsFile(path string) bool {
	fi, e := os.Stat(path)
	if e != nil {
		return false //不存在
	}
	return !fi.IsDir()
}
