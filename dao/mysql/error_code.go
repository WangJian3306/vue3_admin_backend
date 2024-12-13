package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
	ErrorInvalidID       = errors.New("无效的ID")

	ErrorTrademarkExist = errors.New("品牌已存在")

	ErrorRoleExist = errors.New("角色已存在")
	ErrorMenuExist = errors.New("菜单已经存在")
)
