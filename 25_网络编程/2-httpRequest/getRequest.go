/*
@ Time : 2020/7/24 11:30
@ Author : qyz
@ File : getRequest
@ Software: GoLand
@ Description:
*/

package main

import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"log"
)

// https://github.com/kirinlabs/HttpRequest
//https://blog.csdn.net/flyfreelyit/article/details/80281467
func main() {

	res, err := HttpRequest.Get("https://api.github.com/events")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(res.StatusCode())
		body, _ := res.Body()
		fmt.Println(string(body))
	}

	res.Close()

}
