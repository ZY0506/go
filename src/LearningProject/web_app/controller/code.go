package controller

type ResCode int

const (
	// CodeSuccess 成功
	CodeSuccess ResCode = 10000 + iota
	// CodeInvalidParam 参数错误
	CodeInvalidParam
	// CodeUserExist 用户已存在
	CodeUserExist
	// CodeUserNotExist 用户不存在
	CodeUserNotExist
	// CodeInvalidPassword 密码错误
	CodeInvalidPassword
	// CodeServerBusy 服务器繁忙
	CodeServerBusy

	// CodeNeedLogin 需要登录
	CodeNeedLogin
	// CodeInvalidToken 无效的Token
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeInvalidToken:    "无效的Token",
	//CodeExpiredToken:     "Token已过期",
	CodeNeedLogin: "需要登录",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
