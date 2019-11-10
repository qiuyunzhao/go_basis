package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	//字段名称要大写，否则获取不到
	Name string
	Age  int
}

func main() {
	//请求路径 http://127.0.0.1:8081
	http.HandleFunc("/", Welcome)

	//使文件服务生效
	//当发现url以 "/static" 开头时，把请求转发给指定路径
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("26_GoWeb/3_html模板和静态资源/static"))))

	http.ListenAndServe("127.0.0.1:8081", nil)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("26_GoWeb/3_html模板和静态资源/view/index.html")
	if err != nil {
		log.Println(err.Error())
	}
	//往html传入参数
	//t.Execute(w, "qyz") //第二个参数为 传给html数据（html中用{{.}}获取）

	//往html传入结构体
	t.Execute(w, User{Name: "张三", Age: 18})
}
