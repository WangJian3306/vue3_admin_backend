package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"vue3_admin/dao/mysql"
	"vue3_admin/model"
	"vue3_admin/pkg/translation"
	"vue3_admin/service"
)

var UserController userController

type userController struct {
}

// Login 处理登录请求的函数
// @Summary 用户登录接口
// @Description 处理用户登录请求
// @Tags 后台登录与菜单管理
// @Accept application/json
// @Produce application/json
// @Param object body model.ParamUserLogin true "用户登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseLogin
// @Router /admin/acl/index/login [post]
func (u *userController) Login(c *gin.Context) {
	// 1. 获取参数及参数校验
	p := new(model.ParamUserLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2. 业务处理
	token, err := service.UserService.Login(p)
	if err != nil {
		zap.L().Error("Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)

		return
	}
	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"code":    CodeSuccess, // id 值大于 1 <<53-1 int64类型的最大值是1<<63-1
		"message": "success",
		"ok":      true,
		"data":    token,
	})

}

// SignUp 处理新增用户接口
// @Summary 用户新增接口
// @Description 处理用户新增请求
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param object body model.ParamUserSignUp true "用户登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseUserSingUP
// @Router /admin/acl/user/save [post]
func (u *userController) SignUp(c *gin.Context) {
	// 1. 参数校验
	p := new(model.ParamUserSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断 err 是不是 validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非 validator.validationErrors 类型错误直接返回
			ResponseError(c, CodeInvalidParam)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(
			c,
			CodeInvalidParam,
			translation.RemoveTopStruct(errs.Translate(translation.Trans)), // 翻译错误
		)
		return
	}
	// 手动对请求参数进行详细的业务规则校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	//	zap.L().Error("SignUp with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//}

	// 2. 业务处理
	if err := service.UserService.SignUp(p); err != nil {
		zap.L().Error("SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// GetInfo 获取用户信息处理函数
// @Summary 用户信息接口
// @Description 获取用户信息
// @Tags 后台登录与菜单管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseUserInfo
// @Router /admin/acl/index/info [get]
func (u *userController) GetInfo(c *gin.Context) {
	// 1. 获取当前用户ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	// 2. 获取用户信息
	userInfo, err := service.UserService.GetUserInfo(userID)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, userInfo)
}

// Logout 处理登出函数
// @Summary 用户登出接口
// @Description 处理用户登出
// @Tags 后台登录与菜单管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/acl/index/logout [post]
func (*userController) Logout(c *gin.Context) {
	ResponseSuccess(c, nil)
}

// GetUser 获取用户列表
// @Summary 获取用户分页列表
// @Description 获取用户列表
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param page path int true "当前页码"
// @Param limit path int true "每页记录数"
// @Param username query string false "用户名"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseUserList
// @Router /admin/acl/user/{page}/{limit} [get]
func (*userController) GetUser(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	username := c.Query("username")

	// 获取数据
	data, err := service.UserService.GetUserList(username, page, size)
	if err != nil {
		zap.L().Error("service.UserService.GetUserList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 返回响应
	ResponseSuccess(c, data)
}

// UpdateUser 更新用户
// @Summary 更新用户接口
// @Description 处理更新用户请求
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param object body model.ParamUserUpdate true "用户登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/acl/user/update [put]
func (*userController) UpdateUser(c *gin.Context) {
	p := new(model.ParamUserUpdate)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateUser with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err := service.UserService.UpdateUser(p)
	if err != nil {
		zap.L().Error("UpdateUser failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// DeleteUser 删除用户
// @Summary 删除用户接口
// @Description 处理删除用户请求
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param id path int true "用户 ID"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/acl/user/remove/{id} [delete]
func (*userController) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	userId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("DeleteUser with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	err = service.UserService.DeleteUser(userId)
	if err != nil {
		zap.L().Error("service.UserService.DeleteUser() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// BatchDeleteUser 批量删除用户
// @Summary 批量删除用户接口
// @Description 处理批量删除用户请求
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param object body []int64 true "用户 ID 列表"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/acl/user/batchRemove [delete]
func (*userController) BatchDeleteUser(c *gin.Context) {
	var userIDs []int64
	// 绑定 JSON 请求体到 userIDs 切片
	if err := c.ShouldBindJSON(&userIDs); err != nil {
		zap.L().Error("BatchDeleteUser with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 检查用户 ID 列表是否为空
	if len(userIDs) == 0 {
		zap.L().Error("BatchDeleteUser with empty userIDs")
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 调用服务层删除用户
	err := service.UserService.BatchDeleteUser(userIDs)
	if err != nil {
		zap.L().Error("service.UserService.BatchDeleteUser() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// ToAssign 获取用户角色分配接口
// @Summary 用户角色分配接口
// @Description 获取用户角色分配息
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param adminId path int true "用户 ID"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseToAssignRoleList
// @Router /admin/acl/user/toAssign/{adminId} [get]
func (*userController) ToAssign(c *gin.Context) {
	idStr := c.Param("adminId")
	userId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get role detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := service.UserService.ToAssign(userId)
	if err != nil {
		zap.L().Error("service.UserService.ToAssign() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// DoAssignRole 为用户分配角色
// @Summary 为用户分配角色
// @Description 为用户分配角色
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param object body model.ParamDoAssignRole true "用户角色"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/acl/user/doAssignRole [post]
func (*userController) DoAssignRole(c *gin.Context) {
	p := new(model.ParamDoAssignRole)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("DoAssignRole with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err := service.UserService.DoAssignRole(p)
	if err != nil {
		zap.L().Error("service.UserService.DoAssignRole() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
