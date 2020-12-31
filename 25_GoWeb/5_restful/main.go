package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//restful只是一种url风格
func main() {
	r := mux.NewRouter()
	//http://127.0.0.1:8081/qyz/abc Get方式
	r.HandleFunc("/qyz/{url}", func1).Methods("GET")
	http.ListenAndServe(":8081", r)
}

func func1(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, vars["url"])
}
