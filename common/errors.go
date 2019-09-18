package common

import (
	"fmt"
)

type Err struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (err Err) Error() string {
	return fmt.Sprintf("ErrCode:%d , ErrMsg:%s", err.Code, err.Msg)
}

func (err Err) SetError(code int, msg string) Err {
	return Err{code, msg}
}

func (err Err) GetError() string {
	return fmt.Sprintf("ErrCode:%d , ErrMsg:%s", err.Code, err.Msg)
}

var (
	Success         = Err{Code: 200, Msg: "SUCCESS"}
	ErrClientParams = Err{Code: 4000, Msg: "缺少必要参数，或者参数值/路径格式不正确。"}
	ErrUnKnow       = Err{Code: 9000, Msg: "未知错误，请稍后重试。"}
)

var (
	ErrUserEmpty = Err{Code: 100100, Msg: "用户不存在。"}
	ErrPwdEmpty  = Err{Code: 100100, Msg: "密码不能为空。"}
	ErrUserOrPwd = Err{Code: 100100, Msg: "帐号或密码错误。"}
)

var (
	ErrUserExist       = Err{Code: 100100, Msg: "用户已存在,请修改后重试。"}
	ErrUserLogin       = Err{Code: 100101, Msg: "用户名或密码错误,请检查后重试。"}
	ErrDataUnExist     = Err{Code: 100200, Msg: "数据信息不存在,请检查后重试。"}
	ErrDataCreate      = Err{Code: 100201, Msg: "数据信息插入失败,请检查后重试。"}
	ErrDataUpdate      = Err{Code: 100202, Msg: "数据信息更新失败,请检查后重试。"}
	ErrDataGet         = Err{Code: 100203, Msg: "数据信息获取失败,请检查后重试。"}
	ErrDataFind        = Err{Code: 100204, Msg: "数据信息获取失败,请检查后重试。"}
	ErrDataEmpty       = Err{Code: 100205, Msg: "数据信息不存在,请检查后重试。"}
	ErrDataEmptyParams = Err{Code: 100206, Msg: "缺少必要参数,请检查后重试。"}
	ErrDataNoExist     = Err{Code: 10030, Msg: "数据信息不存在存在,请检查后重试。"}
	ErrLuckFinal       = Err{Code: 10040, Msg: "奖品已发放完毕。"}
	ErrLuckFail        = Err{Code: 10050, Msg: "活动太火爆了，请稍后重试。"}
)
