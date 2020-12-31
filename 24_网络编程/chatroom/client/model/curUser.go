package model

import (
	"go_basis/24_网络编程/chatroom/common/message"
	"net"
)

//因为在客户端，我们很多地方会使用到curUser,我们将其作为一个全局
type CurUser struct {
	Conn net.Conn
	message.User
}
