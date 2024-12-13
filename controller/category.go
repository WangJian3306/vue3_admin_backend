package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"vue3_admin/model"
	"vue3_admin/service"
)

var CategoryController categoryController

type categoryController struct {
}

// GetCategory1 获取一级分类
// @Summary 获取一级分类接口
// @Description 获取一级分类
// @Tags 商品分类接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseTmList
// @Router /admin/product/getCategory1 [get]
func (*categoryController) GetCategory1(c *gin.Context) {
	data, err := service.CategoryService.GetCategory1()
	if err != nil {
		zap.L().Error("service.CategoryService.GetCategory1() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

// GetCategory2 获取二级分类
// @Summary 获取二级分类接口
// @Description 获取二级分类
// @Tags 商品分类接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param id path int true "分类一 ID"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseTmList
// @Router /admin/product/getCategory2/{id} [get]
func (*categoryController) GetCategory2(c *gin.Context) {
	idStr := c.Param("id")
	category1Id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := service.CategoryService.GetCategory2(category1Id)
	if err != nil {
		zap.L().Error("service.CategoryService.GetCategory2() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

// GetCategory3 获取三级分类
// @Summary 获取三级分类接口
// @Description 获取三级分类
// @Tags 商品分类接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param id path int true "分类二 ID"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseTmList
// @Router /admin/product/getCategory3/{id} [get]
func (*categoryController) GetCategory3(c *gin.Context) {
	idStr := c.Param("id")
	category2Id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := service.CategoryService.GetCategory3(category2Id)
	if err != nil {
		zap.L().Error("service.CategoryService.GetCategory3() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

func (*categoryController) CreateCategory2(c *gin.Context) {
	p := new(model.ParamC2Create)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("CreateCategory2 with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 业务处理
	err := service.CategoryService.CreateCategory2(p)
	if err != nil {
		zap.L().Error("service.TrademarkService.CreateTrademark() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回结果
	ResponseSuccess(c, nil)
}

func (*categoryController) CreateCategory3(c *gin.Context) {
	p := new(model.ParamC3Create)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 业务处理
	err := service.CategoryService.CreateCategory3(p)
	if err != nil {
		zap.L().Error("service.TrademarkService.CreateTrademark() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回结果
	ResponseSuccess(c, nil)
}
