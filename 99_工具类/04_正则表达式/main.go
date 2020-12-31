package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("------------------------------------- 简单匹配 -----------------------------------------")
	res, _ := regexp.MatchString(`\w`, "abc123") //匹配非特殊字符，即a-z、A-Z、0-9、_、汉字
	fmt.Println(res)
	res, _ = regexp.MatchString(`\d`, "abc/'[1;-") //匹配数字，即0-9
	fmt.Println(res)
	res, _ = regexp.MatchString(`\D`, "562k6456") //匹配非数字，即不是数字
	fmt.Println(res)

	fmt.Println("------------------------------------- 复杂匹配 -----------------------------------------")
	pattern1 := `\d[a-z0-9]` //数字+小写字母或数字
	reg := regexp.MustCompile(pattern1)
	fmt.Println(reg.MatchString("hhh1gdu123ji88111"))

	fmt.Println(reg.FindAllString("hhh1gdu123ji88", 7))        //返回n个符合要求的子串   n为-1表示返回所有符合要求子串
	fmt.Println(reg.Split("hhh1gdu123ji88", -1))               //根据符合的子串去切割
	fmt.Println(reg.ReplaceAllString("hhh1gdu123ji88", "替换值")) //把符合的子串代替为 repl 的值

	fmt.Println("------------------------------------- 验证邮箱 -----------------------------------------")
	fmt.Println(VerifyEmailFormat("12345@126.com"))         //true
	fmt.Println(VerifyEmailFormat("1694644469@qq.com"))     //true
	fmt.Println(VerifyEmailFormat("qiuyunzhao@inspur.com")) //true
	fmt.Println(VerifyEmailFormat("12345126.com"))          //false
}

//匹配电子邮箱
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
