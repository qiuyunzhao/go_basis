package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	////使文件服务生效(当发现url以 "/static" 开头时，把请求转发给指定路径)
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("26_GoWeb/4_Cookie/static"))))

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

	t, err := template.ParseFiles("26_GoWeb/4_Cookie/view/index.html")
	if err != nil {
		log.Println(err.Error())
	}
	t.Execute(w, cookies) //将cookies传去html
}

//访问服务器时，服务器端产生cookie
func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "qyz", Value: "myValue"}
	http.SetCookie(w, &cookie)
	Welcome(w, r)
}

//访问服务器时，服务器端不产生cookie
func Welcome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("26_GoWeb/4_Cookie/view/index.html")
	if err != nil {
		log.Println(err.Error())
	}
	t.Execute(w, nil)
}
