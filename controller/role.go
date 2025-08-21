package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"vue3_admin/model"
	"vue3_admin/service"
)

var RoleController roleController

type roleController struct {
}

// GetRole 获取角色列表
// @Summary 获取角色分页列表
// @Description 获取角色列表
// @Tags 角色管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param page path int true "当前页码"
// @Param limit path int true "每页记录数"
// @Param roleName query string false "角色名"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseRoleList
// @Router /admin/acl/role/{page}/{limit} [get]
func (r *roleController) GetRole(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	roleName := c.Query("roleName")

	// 获取数据
	data, err := service.RoleService.GetRoleList(roleName, page, size)
	if err != nil {
		zap.L().Error("service.RoleService.GetRoleList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回数据
	ResponseSuccess(c, data)
}

// SaveRole 新增角色
// @Summary 新增角色接口
// @Description 处理新增角色请求
// @Tags 角色管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param object body model.ParamRoleSave true "用户登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/acl/role/save [post]
func (r *roleController) SaveRole(c *gin.Context) {
	p := new(model.ParamRoleSave)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SaveRole with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	err := service.RoleService.SaveRole(p)
	if err != nil {
		zap.L().Error("service.RoleService.SaveRole() failed", zap.String("role_name", p.RoleName), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// UpdateRole 更新角色
// @Summary 更新角色接口
// @Description 处理更新角色请求
// @Tags 角色管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param object body model.ParamRoleUpdate true "用户登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/acl/role/update [put]
func (r *roleController) UpdateRole(c *gin.Context) {
	p := new(model.ParamRoleUpdate)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateRole with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err := service.RoleService.UpdateRole(p)
	if err != nil {
		zap.L().Error("service.RoleService.UpdateRole() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// DeleteRole 删除角色
// @Summary 删除角色接口
// @Description 处理删除角色请求
// @Tags 角色管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param id path int true "角色 ID"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/acl/role/remove/{id} [delete]
func (r *roleController) DeleteRole(c *gin.Context) {
	idStr := c.Param("id")
	roleId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("DeleteRole with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err = service.RoleService.DeleteRole(roleId)
	if err != nil {
		zap.L().Error("service.RoleService.DeleteRole() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
