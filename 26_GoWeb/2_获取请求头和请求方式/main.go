package main

import (
	"fmt"
	"net/http"
)

//get 请求：localhost:8081/param?name=zhangSan&age=18
func main() {
	http.HandleFunc("/param", Parm)
	http.ListenAndServe("127.0.0.1:8081", nil)
}

func Parm(w http.ResponseWriter, r *http.Request) {
	//获取请求头
	h := r.Header
	w.Header().Set("Content-Type", "text/html;charset=utf-8") //设置响应头
	fmt.Fprintln(w, h)
	fmt.Fprintln(w, h.Get("User-Agent"))

	//获取请求体
	r.ParseForm() //需要先解析form才能获取到内容
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.Form["name"][0])
	fmt.Fprintln(w, r.FormValue("age"))

}
