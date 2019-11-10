package main

import (
	"github.com/gorilla/websocket"
	"go_code/30_WebSocket/ws"
	"net/http"
)

// http升级为websocket协议的配置
var wsUpgrader = websocket.Upgrader{
	// 回调函数CheckOrigin：允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	//请求： http://127.0.0.1:7777/ws
	http.HandleFunc("/ws", WebSocketHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}

func WebSocketHandler(resp http.ResponseWriter, req *http.Request) {
	// 应答客户端告知升级连接为websocket
	wsSocket, err := wsUpgrader.Upgrade(resp, req, nil)
	if err != nil {
		return
	}

	wsConn, _ := ws.InitWSConnection(wsSocket)

	// 处理器
	go wsConn.ProcLoop()
	// 读WebSocket协程
	go wsConn.WsReadLoop()
	// 写WebSocket协程
	go wsConn.WsWriteLoop()
}
