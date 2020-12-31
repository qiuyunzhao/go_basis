/*
@ Time : 2020/7/24 11:30
@ Author : qyz
@ File : getRequest
@ Software: GoLand
@ Description:
*/

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"io/ioutil"
	"log"
	"net/http"
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

/**
 * @ Time:  2020/7/24 9:15
 * @ Description: 获取所有grafana文件夹 -- GET请求
 * @ Param:
 *         url : 请求地址
 * @ return:
 *         []GrafanaFolder : 文件夹信息数组
**/
type GetGrafanaFolderData struct {
	Id    int    `json:"id"`
	Uid   string `json:"uid"`
	Title string `json:"title"`
}

func getGrafanaFolders(url string) ([]GetGrafanaFolderData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.Status != "200 OK" {
		return nil, errors.New("grafana请求失败,响应码：" + resp.Status)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var folders []GetGrafanaFolderData
	if err := json.Unmarshal(body, &folders); err != nil {
		return nil, err
	}

	return folders, nil
}

/**
 * @ Time:  2020/7/24 10:53
 * @ Description: 创建grafana文件夹 -- POST请求
 * @ Param:
 *         url : 请求地址
 *         folderMsg : 文件夹信息
 * @ return:
 *          :
**/
type CreateGrafanaFolderData struct {
	Uid   string `json:"uid"`
	Title string `json:"title"`
}

func CreateGrafanaFolder(url string, folderMsg CreateGrafanaFolderData) error {
	jsonData, err := json.Marshal(folderMsg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.Status != "200 OK" {
		return errors.New("grafana请求失败,响应码：" + resp.Status)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}

/**
 * @ Time:  2020/7/24 11:05
 * @ Description: 更新grafana文件夹 -- PUT请求
 * @ Param:
 *         url : 请求地址
 *         folderMsg : 文件夹信息
 * @ return:
 *          :
**/
type UpdateGrafanaFolderData struct {
	Uid       string `json:"uid"`
	Title     string `json:"title"`
	Overwrite bool   `json:"overwrite"`
}

func UpdateGrafanaFolderByUid(url string, folderMsg UpdateGrafanaFolderData) error {
	jsonData, err := json.Marshal(folderMsg)
	if err != nil {
		return err
	}

	url = url + "/" + folderMsg.Uid
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.Status != "200 OK" {
		return errors.New("grafana请求失败,响应码：" + resp.Status)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

/**
 * @ Time:  2020/7/24 11:05
 * @ Description: 删除grafana文件夹 -- DELETE请求
 * @ Param:
 *         url : 请求地址
 *         folderMsg : 文件夹编号
 * @ return:
 *          :
**/
func DeleteGrafanaFolderByUid(url string, uid string) error {
	url = url + "/" + uid
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.Status != "200 OK" {
		return errors.New("grafana请求失败,响应码：" + resp.Status)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}
