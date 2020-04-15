package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("-------------------------------------------  1  -------------------------------------------------")
	//统计字符串的长度，按字节 len(str)
	//golang的编码统一为utf-8 (ascii的字符(字母和数字) 占一个字节，汉字占用3个字节)
	str := "hello北"
	fmt.Println("str len=", len(str)) // 8

	fmt.Println("-------------------------------------------  2  -------------------------------------------------")
	str2 := "hello北京"
	//字符串遍历，同时处理有中文的问题 r := []rune(str) 切片
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c \n", r[i])
	}

	fmt.Println("-------------------------------------------  3  -------------------------------------------------")
	//字符串转整数:
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转成的结果是", n)
	}

	fmt.Println("-------------------------------------------  4  -------------------------------------------------")
	//4)整数转字符串:
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v, str=%T \n", str, str)

	fmt.Println("-------------------------------------------  5  -------------------------------------------------")
	//5)字符串 转 []byte:
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v \n", bytes)

	fmt.Println("-------------------------------------------  6  -------------------------------------------------")
	//6)[]byte 转 字符串:
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v \n", str)

	fmt.Println("-------------------------------------------  7  -------------------------------------------------")
	//10进制转 2, 8, 16进制:
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v \n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的16进制是=%v \n", str)

	fmt.Println("-------------------------------------------  8  -------------------------------------------------")
	//查找子串是否在指定的字符串中:
	b := strings.Contains("seafood", "mary")
	fmt.Printf("b=%v \n", b)
	b = strings.Contains("hello,world!", "world")
	fmt.Printf("b=%v \n", b)

	fmt.Println("-------------------------------------------  9  -------------------------------------------------")
	//统计一个字符串有几个指定的子串:
	num := strings.Count("ceheese", "e")
	fmt.Printf("num=%v \n", num)

	fmt.Println("------------------------------------------  10  -------------------------------------------------")
	//10)不区分大小写的字符串比较(==是区分字母大小写的):
	b = strings.EqualFold("abc", "Abc") //不区分大小写
	fmt.Printf("b=%v \n", b)            //true
	fmt.Println("结果", "abc" == "Abc")   // false ; "==" 区分字母大小写

	fmt.Println("------------------------------------------  11  -------------------------------------------------")
	//11)返回子串在字符串第一次出现的index值（0开始），如果没有返回-1 :
	index := strings.Index("NLT_abcabcabc", "abc") // 4
	fmt.Printf("index=%v \n", index)

	fmt.Println("------------------------------------------  12  -------------------------------------------------")
	//12)返回子串在字符串最后一次出现的index（0开始），如没有返回-1
	index = strings.LastIndex("go golang", "go") //3
	fmt.Printf("index=%v \n", index)

	fmt.Println("------------------------------------------  13  -------------------------------------------------")
	//将指定的子串替换成 另外一个子串: strings.Replace("go go hello", "go", "go语言", n)
	//n可以指定你希望替换几个(可以是变量)，如果n=-1表示全部替换
	str2 = "go go hello"
	str = strings.Replace(str2, "go", "北京", -1)
	fmt.Printf("str2=%v \n", str2)
	fmt.Printf("str=%v \n", str)

	fmt.Println("------------------------------------------  14  -------------------------------------------------")
	//按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：
	strArr := strings.Split("hello,wrold,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v \n", i, strArr[i])
	}
	fmt.Printf("strArr=%v \n", strArr)

	fmt.Println("------------------------------------------  15  -------------------------------------------------")
	//15)将字符串的字母进行大小写的转换:
	str = "goLang Hello"
	str = strings.ToLower(str)
	fmt.Printf("str=%v \n", str) //golang hello
	str = strings.ToUpper(str)
	fmt.Printf("str=%v \n", str) //GOLANG HELLO

	fmt.Println("------------------------------------------  16  -------------------------------------------------")
	//将字符串左右两边的空格去掉:
	str = strings.TrimSpace(" tn a lone gopher ntrn   ")
	fmt.Printf("str=%q \n", str) //str="tn a lone gopher ntrn"

	//将字符串左右两边指定的字符去掉 ：
	str = strings.Trim("! he!llo! ", " !") //将左右两边 ! 和 " "去掉
	fmt.Printf("str=%q \n", str)

	//将字符串左两边指定的字符去掉 ：
	str = strings.TrimLeft("! he!llo! ", " !") //将左右两边 ! 和 " "去掉
	fmt.Printf("str=%q \n", str)

	//将字符串右两边指定的字符去掉 ：
	str = strings.TrimRight("! he!llo! ", " !") //将左右两边 ! 和 " "去掉
	fmt.Printf("str=%q \n", str)

	fmt.Println("------------------------------------------  17  -------------------------------------------------")
	//判断字符串是否以指定的字符串开头:
	b = strings.HasPrefix("ftp://192.168.10.1", "hsp") //true
	fmt.Printf("b=%v \n", b)

}
