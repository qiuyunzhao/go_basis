package main

import (
	"fmt"
	"regexp"
)

func main() {
	res, _ := regexp.MatchString(`\w`, "abc123") //匹配非特殊字符，即a-z、A-Z、0-9、_、汉字
	fmt.Println(res)
	res, _ = regexp.MatchString(`\d`, "abc/'[1;-") //匹配数字，即0-9
	fmt.Println(res)
	res, _ = regexp.MatchString(`\D`, "562k6456") //匹配非数字，即不是数字
	fmt.Println(res)

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
