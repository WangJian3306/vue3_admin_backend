package service

import (
	"go.uber.org/zap"
	"vue3_admin/dao/mysql"
	"vue3_admin/model"
	"vue3_admin/pkg/snowflake"
)

var RoleService roleService

type roleService struct {
}

func (r *roleService) GetRoleList(roleName string, page, limit int64) (data *model.ResponseRoleList, err error) {
	data, err = mysql.RoleDao.GetRoleList(roleName, page, limit)
	if err != nil {
		zap.L().Error("mysql.UserDao.GetUserList() failed", zap.Error(err))
		return
	}
	return
}

func (r *roleService) SaveRole(p *model.ParamRoleSave) (err error) {
	// 1. 判断角色是否已经存在
	if err = mysql.RoleDao.CheckRoleExist(p.RoleName); err != nil {
		return err
	}

	roleId := snowflake.GenID()
	// 2. 生成一个角色实例
	role := &model.Role{
		RoleID:   roleId,
		RoleName: p.RoleName,
		Remark:   p.Remark,
	}

	// 3. 保存数据库
	err = mysql.RoleDao.InsertRole(role)

	return err
}

func (r *roleService) UpdateRole(p *model.ParamRoleUpdate) (err error) {
	// 构造一个角色实例
	role := &model.Role{
		RoleID:   p.RoleID,
		RoleName: p.RoleName,
		Remark:   p.Remark,
	}

	// 保存进数据库
	err = mysql.RoleDao.UpdateRole(role)

	return err
}

func (r *roleService) DeleteRole(roleId int64) (err error) {
	return mysql.RoleDao.DeleteRole(roleId)
}
