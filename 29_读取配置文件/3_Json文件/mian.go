package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

func fileGetContents(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func main() {
	var c Contact

	content, err := fileGetContents("29_读取配置文件/3_Json文件/data.json")
	if err != nil {
		fmt.Println("open file error: " + err.Error())
		return
	}
	err = json.Unmarshal(content, &c)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return
	}
	fmt.Println(c)
}
