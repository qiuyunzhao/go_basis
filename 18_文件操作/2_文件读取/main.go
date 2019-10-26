package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//test01()	//带缓冲的文件读取-适用于大文件
	test02() //不带缓冲的文件读取-适用于小文件
}

//带缓冲的文件读取-适用于大文件，需要先打开文件 bufio.NewReader(file)
func test01() {
	//打开文件
	file, err := os.Open("go_code/18_文件操作/sources/a.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//当函数退出时，要及时的关闭file，否则会有内存泄漏
	defer file.Close()

	// 创建一个 *Reader ,默认的缓冲区为4096 适用于大文件
	reader := bufio.NewReader(file)

	//循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF {                  // io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Printf(str)
	}

	fmt.Println("文件读取结束...")
}

//不带缓冲的文件读取（一次性读取整个文件）-适用于小文件 不需要先打开文件
func test02() {
	filePath := "go_code/18_文件操作/sources/a.txt"
	content, err := ioutil.ReadFile(filePath) //使用ioutil.ReadFile(file)一次性将文件读取
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}

	fmt.Printf("%v", string(content)) // content->[]byte

	//我们没有显式的Open文件，因此也不需要显式的Close文件
	//因为,文件的Open和Close被封装到 ReadFile 函数内部
}
