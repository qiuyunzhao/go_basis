package main

import (
	"fmt"
	"os"
)

func main() {
	file := "18_文件操作/sources/赵敏.jpg"
	fmt.Println(IsExist(file))
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
