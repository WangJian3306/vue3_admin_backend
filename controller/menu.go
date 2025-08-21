package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"vue3_admin/model"
	"vue3_admin/service"
)

var MenuController menuController

type menuController struct{}

// GetMenu 获取菜单列表
// @Summary 获取菜单列表
// @Description 获取菜单列表
// @Tags 菜单管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseMenuList
// @Router /admin/acl/permission [get]
func (*menuController) GetMenu(c *gin.Context) {
	data, err := service.MenuService.GetMenu()
	if err != nil {
		zap.L().Error("service.MenuService.GetMenu() failed", zap.Error(err))
	}

	ResponseSuccess(c, data)
}

// SaveMenu 新增菜单
// @Summary 新增菜单接口
// @Description 处理新增菜单请求
// @Tags 菜单管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param object body model.ParamMenuSave true "用户登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/acl/permission/save [post]
func (*menuController) SaveMenu(c *gin.Context) {
	p := new(model.ParamMenuSave)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SaveMenu with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err := service.MenuService.SaveMenu(p)
	if err != nil {
		zap.L().Error("service.MenuService.SaveMenu() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// UpdateMenu 更新菜单
// @Summary 更新菜单接口
// @Description 处理更新菜单请求
// @Tags 菜单管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param object body model.ParamMenuUpdate true "用户登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/acl/permission/update [put]
func (*menuController) UpdateMenu(c *gin.Context) {
	p := new(model.ParamMenuUpdate)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateMenu with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err := service.MenuService.UpdateMenu(p)
	if err != nil {
		zap.L().Error("service.MenuService.UpdateMenu() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// DeleteMenu 删除菜单
// @Summary 删除菜单接口
// @Description 处理删除菜单请求
// @Tags 菜单管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param id path int true "菜单 ID"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/acl/permission/remove/{id} [delete]
func (*menuController) DeleteMenu(c *gin.Context) {
	idStr := c.Param("id")
	menuId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("DeleteMenu with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err = service.MenuService.DeleteMenu(menuId)
	if errors.Is(err, model.CodeMenuNodeExistError) {
		zap.L().Error("service.MenuService.DeleteMenu() failed", zap.Error(err))
		ResponseError(c, CodeMenuNodeExist)
		return
	}
	if err != nil {
		zap.L().Error("service.MenuService.DeleteMenu() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}

// ToAssign 根据角色获取菜单
// @Summary 根据角色获取菜单接口
// @Description 根据角色获取菜单请求
// @Tags 菜单管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param roleId path int true "角色ID"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseToAssignMenuList
// @Router /admin/acl/permission/toAssign/{roleId} [get]
func (*menuController) ToAssign(c *gin.Context) {
	idStr := c.Param("roleId")
	fmt.Println(idStr)
	roleId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("ToAssign with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := service.MenuService.ToAssign(roleId)
	if err != nil {
		zap.L().Error("service.MenuService.ToAssign() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// DoAssign 给角色分配权限
// @Summary 给角色分配权限
// @Description 给角色分配权限
// @Tags 菜单管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param roleId query int true "角色 ID"
// @Param permissionId query []int true "菜单 ID列表"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/acl/permission/doAssign [post]
func (*menuController) DoAssign(c *gin.Context) {
	idStr := c.Query("roleId")
	roleId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("DoAssign with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	permissionIdStr := c.Query("permissionId")
	strArr := strings.Split(permissionIdStr, ",")
	var intArr []int64
	for _, s := range strArr {
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			zap.L().Error("permissionId convert failed", zap.Error(err))
			ResponseError(c, CodeInvalidParam)
			return
		}
		intArr = append(intArr, num)
	}

	err = service.MenuService.DoAssign(roleId, intArr)
	if err != nil {
		zap.L().Error("service.MenuService.DoAssign() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}
