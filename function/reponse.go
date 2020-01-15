package function

import (
	"go-common/constants"
)

type BaseResponseInfo struct {
	// 状态码
	Code int64 `json:"code"`

	// 数据结构体
	Data interface{} `json:"data"`

	// 数据信息
	Msg string `json:"msg"`
}

func FunJsonSuccess(data interface{}, msg string) BaseResponseInfo {
	if msg == "" {
		msg = constants.ResponseMessageSuccessDefault
	}
	response := BaseResponseInfo{Code: constants.ResponseCodeSuccess, Data: data, Msg: msg}
	return response
}

func FunJsonSuccessNormal(data interface{}) BaseResponseInfo {
	msg := constants.ResponseMessageSuccessDefault
	response := BaseResponseInfo{Code: constants.ResponseCodeSuccess, Data: data, Msg: msg}
	return response
}

func FunJsonError(code int64, data interface{}, msg string) BaseResponseInfo {
	if msg == "" {
		msg = constants.ResponseMessageErrorNormalDefault
	}
	response := BaseResponseInfo{Code: code, Data: data, Msg: msg}
	return response
}

/**
 * 一般错误返回
 * @param msg string
 */
func FunJsonErrorNormal(msg string) BaseResponseInfo {
	return FunJsonError(constants.ResponseCodeErrorNormal, nil, msg)
}

