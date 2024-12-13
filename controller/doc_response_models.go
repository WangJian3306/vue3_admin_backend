package controller

import (
	"vue3_admin/model"
)

// 专门用来放接口文档用到的 model
// 因为我们的接口文档返回的数据格式是一致的，但是具体的 data 类型不一致

type _BaseResponse struct {
	Code    ResCode `json:"code"`    // 业务响应状态码
	Message string  `json:"message"` // 提示信息
	OK      bool    `json:"ok"`
}

type _ResponseLogin struct {
	Code    ResCode `json:"code"`    // 业务响应状态码
	Message string  `json:"message"` // 提示信息
	Data    string  `json:"data"`    // 响应数据
	OK      bool    `json:"ok"`
}

type _ResponseUserInfo struct {
	_BaseResponse
	Data model.ResponseUserInfo `json:"data"`
}

type _ResponseUserSingUP struct {
	_BaseResponse
	Data string `json:"data"`
}

type _ResponseUserList struct {
	_BaseResponse
	Data model.ResponseUserList `json:"data"`
}

type _ResponseRoleList struct {
	_BaseResponse
	Data model.ResponseRoleList `json:"data"`
}

type _ResponseToAssignRoleList struct {
	_BaseResponse
	Data model.ResponseToAssignRole `json:"data"`
}

type _ResponseMenuList struct {
	_BaseResponse
	Data model.Menu `json:"data"`
}

type _ResponseToAssignMenuList struct {
	_BaseResponse
	Data []model.Menu `json:"data"`
}

type _ResponseTmList struct {
	_BaseResponse
	Data model.ResponseTmList `json:"data"`
}
