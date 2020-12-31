package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	////使文件服务生效(当发现url以 "/static" 开头时，把请求转发给指定路径)
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("25_GoWeb/4_Cookie/static"))))

	http.HandleFunc("/", Welcome)
	http.HandleFunc("/noCookie", Welcome)
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	http.ListenAndServe("127.0.0.1:8081", nil)
}

//获取cookie
func getCookie(w http.ResponseWriter, r *http.Request) {
	//r.Cookie("quz") //根据Name取出Cookie
	cookies := r.Cookies() //取出全部Cookie

	t, err := template.ParseFiles("25_GoWeb/4_Cookie/view/index.html")
	if err != nil {
		log.Println(err.Error())
	}
	t.Execute(w, cookies) //将cookies传去html
}

//访问服务器时，服务器端产生cookie
func setCookie(w http.ResponseWriter, r *http.Request) {
	validTime := time.Now().Add(time.Second * 5)
	cookie := http.Cookie{
		Name:  "qyz",
		Value: "myValue",
		//HttpOnly: true, //设置不能通过js脚本获取
		//Path:     "/abc/", //设置/abc开头 及其子路由可以获取到cookie
		//MaxAge:   3, //设置3s有效，默认关闭浏览器失效（有的浏览器不支持该属性）
		//Expires:time.Date(2019,10,26,19,16,1,1,time.Local), //设置有效期到该时间
		Expires: validTime, //设置有效期5s
	}
	http.SetCookie(w, &cookie)
	Welcome(w, r)
}

//访问服务器时，服务器端不产生cookie
func Welcome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("25_GoWeb/4_Cookie/view/index.html")
	if err != nil {
		log.Println(err.Error())
	}
	t.Execute(w, nil)
}
