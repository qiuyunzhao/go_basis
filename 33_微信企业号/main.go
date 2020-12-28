package main

//# 微信推送配置
//wechat:
//# 企业ID
//corpid: wwd55bf9fc496beac6
//# 应用的凭证密钥
//corpsecret: ygeeu9A10NC0BnFtbs0R1ahTNop4VpBq
//# 应用的id
//agentid: 1000002

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 获取到的Token结构体
type WechatToken struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// 文本卡片消息
type TextCardMsg struct {
	ToUser        string   `json:"touser"`
	ToParty       string   `json:"toparty"`
	ToTag         string   `json:"totag"`
	MsgType       string   `json:"msgtype"`
	AgentId       int      `json:"agentid"`
	TextCard      TextCard `json:"textcard"`
	EnableIdTrans int      `json:"enable_id_trans"`
}

type TextCard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Btntxt      string `json:"btntxt"`
}

func main() {
	wechatToken := getWechatToken("wwd55bf9fc496beac6", "ygeeu9A10NC0BnFtbs0R1ahTNop4VpBq_J4qrn-f24Y")
	fmt.Println("获取的token：", wechatToken)

	textMsg := TextCardMsg{
		ToUser:  "@all",
		ToParty: "",
		ToTag:   "",
		MsgType: "textcard",
		AgentId: 1000002,
		TextCard: TextCard{
			Title: "领奖通知1",
			Description: "<div class=\"gray\">2016年9月26日</div>" +
				"<div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div>" +
				"<div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>",
			Url:    "url",
			Btntxt: "请及时维护！！！",
		},
		EnableIdTrans: 0,
	}
	postJsonMsg(wechatToken, textMsg)

}

//get请求获取Token
func getWechatToken(corpId string, corpSecret string) string {
	//http.Get的参数必须是带http://协议头的完整url,不然请求结果为空
	resp, _ := http.Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpId + "&corpsecret=" + corpSecret)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// 反序列化为结构体
	wechatToken := WechatToken{}
	err := json.Unmarshal(body, &wechatToken)
	if err != nil {
		log.Println(err)
	}

	return wechatToken.AccessToken
}

// post请求，发送Json格式消息
func postJsonMsg(token string, msg interface{}) {
	// 序列化
	body, _ := json.Marshal(msg)

	// post请求提交json数据
	resp, _ := http.Post("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token="+token, "application/json", bytes.NewBuffer(body))

	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Post request with json result: %s\n", string(respBody))
}
