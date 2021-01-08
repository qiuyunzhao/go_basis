/*
@ Time : 2021/1/8 10:04
@ Author : qyz
@ File : code
@ Software: GoLand
@ Description: Gin 响应数据的响应码
*/

package ginResult

// 错误码规则:

// (1) 错误码需为 > 0 的数;
//
// (2) 错误码为 5 位数:
//              ----------------------------------------------------------
//                第1位               2、3位                  4、5位
//              ----------------------------------------------------------
//              服务级错误码         模块级错误码	        具体错误码
//              ----------------------------------------------------------

var (
	// OK
	OK = response(200, "ok")

	// 服务级错误码
	ErrServer    = response(10001, "本服务异常")
	ErrParam     = response(10002, "参数有误")
	ErrSignParam = response(10003, "签名参数有误")

	// 模块级错误码 - 用户模块
	ErrUserService = response(20100, "用户服务异常")
	ErrUserPhone   = response(20101, "用户手机号不合法")
	ErrUserCaptcha = response(20102, "用户验证码有误")

	// 库存模块
	ErrOrderService = response(20200, "订单服务异常")
	ErrOrderOutTime = response(20201, "订单超时")

	// ......
)
