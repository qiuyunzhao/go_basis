package ws

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

// websocket读写消息结构体
type wsMessage struct {
	messageType int
	data        []byte
}

// websocket连接结构体
type wsConnection struct {
	wsSocket  *websocket.Conn // 底层websocket
	readChan  chan *wsMessage // 读队列
	writeChan chan *wsMessage // 写队列
	closeChan chan byte       // 关闭通知
	isClosed  bool            // 关闭标志
	mutex     sync.Mutex      // 避免重复关闭管道
}

// 获取websocket连接
func InitWSConnection(wsSocketConn *websocket.Conn) (conn *wsConnection, err error) {
	conn = &wsConnection{
		wsSocket:  wsSocketConn,
		readChan:  make(chan *wsMessage, 1000),
		writeChan: make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
	}
	return
}

//---------------------------------------------------------------------------------------------------------------------
func (wsConn *wsConnection) WsReadLoop() {
	for {
		// 从WebSocket长连接中 读一个message
		msgType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			goto error
		}
		reqMessage := &wsMessage{
			msgType,
			data,
		}
		// 放入请求队列
		select {
		case wsConn.readChan <- reqMessage:
		case <-wsConn.closeChan:
			goto closed
		}
	}
error:
	wsConn.wsClose()
closed:
}

func (wsConn *wsConnection) WsWriteLoop() {
	for {
		select {
		// 取一个应答
		case msg := <-wsConn.writeChan:
			// 写给websocket
			if err := wsConn.wsSocket.WriteMessage(msg.messageType, msg.data); err != nil {
				goto error
			}
		case <-wsConn.closeChan:
			goto closed
		}
	}
error:
	wsConn.wsClose()
closed:
}

func (wsConn *wsConnection) ProcLoop() {
	// 启动一个gouroutine发送心跳
	go func() {
		for {
			time.Sleep(2 * time.Second)
			if err := wsConn.wsWrite(websocket.TextMessage, []byte("heartbeat from server")); err != nil {
				fmt.Println("heartbeat fail")
				wsConn.wsClose()
				break
			}
		}
	}()

	// 这是一个同步处理模型（只是一个例子），如果希望并行处理可以每个请求一个gorutine，注意控制并发goroutine的数量!!!
	for {
		msg, err := wsConn.wsRead()
		if err != nil {
			fmt.Println("read fail")
			break
		}
		fmt.Println(string(msg.data))
		err = wsConn.wsWrite(msg.messageType, msg.data)
		if err != nil {
			fmt.Println("write fail")
			break
		}
	}
}

//---------------------------------------------------------------------------------------------------------------------
func (wsConn *wsConnection) wsRead() (*wsMessage, error) {
	select {
	case msg := <-wsConn.readChan:
		return msg, nil
	case <-wsConn.closeChan:
	}
	return nil, errors.New("ws closed")
}

func (wsConn *wsConnection) wsWrite(messageType int, data []byte) error {
	select {
	case wsConn.writeChan <- &wsMessage{messageType, data}:
	case <-wsConn.closeChan:
		return errors.New("ws closed")
	}
	return nil
}

func (wsConn *wsConnection) wsClose() {
	wsConn.wsSocket.Close() //线程安全，可重入

	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()

	if !wsConn.isClosed {
		wsConn.isClosed = true
		close(wsConn.closeChan)
	}
}
