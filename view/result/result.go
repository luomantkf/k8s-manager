package result

import "private-cloud-k8s/pkg/e"

//封装统一的返回结果
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func OkResultWithMsg(msg string) Result {
	return Result{
		Code: e.SUCCESS,
		Msg:  msg,
		Data: nil,
	}
}

func OkResultWithData(data interface{}) Result {
	return Result{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
		Data: data,
	}
}

func OkResultWithMsgAndData(msg string, data interface{}) Result {
	return Result{
		Code: e.SUCCESS,
		Msg:  msg,
		Data: data,
	}
}

func FailResultWithMsg(msg string) Result {
	return Result{
		Code: e.FAIL,
		Msg:  msg,
		Data: nil,
	}
}

func FailResultWithData(data interface{}) Result {
	return Result{
		Code: e.FAIL,
		Msg:  e.GetMsg(e.FAIL),
		Data: data,
	}
}

func FailResultWithMsgAndData(msg string, data interface{}) Result {
	return Result{
		Code: e.FAIL,
		Msg:  msg,
		Data: data,
	}
}
