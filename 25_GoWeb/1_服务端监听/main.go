package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Welcome)
	http.HandleFunc("/abc", Welcome1)
	http.ListenAndServe("127.0.0.1:8081", nil)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "服务器收到请求，并返回信息：<b>你好，欢迎访问！</b>")
}

func Welcome1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8") //设置响应头
	fmt.Fprintln(w, "服务器收到请求，并返回信息：<b>来了老弟！</b>")
}
