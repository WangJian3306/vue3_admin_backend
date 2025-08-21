package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"vue3_admin/model"
	"vue3_admin/service"
)

var TrademarkController trademarkController

type trademarkController struct {
}

// CreateTrademark 新增品牌接口
// @Summary 新增品牌接口
// @Description 新增品牌接口
// @Tags 品牌管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param object body model.ParamTmSave true "用户角色"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/product/baseTrademark/save [post]
func (*trademarkController) CreateTrademark(c *gin.Context) {
	// 1. 获取参数及参数校验
	p := new(model.ParamTmSave)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("CreateTrademark with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 业务处理
	err := service.TrademarkService.CreateTrademark(p)
	if err != nil {
		zap.L().Error("service.TrademarkService.CreateTrademark() failed", zap.String("tm_name", p.TmName), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回结果
	ResponseSuccess(c, nil)
}

// GetTrademark 获取品牌列表
// @Summary 获取品牌分页列表
// @Description 获取品牌列表
// @Tags 品牌管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param page path int true "当前页码"
// @Param limit path int true "每页记录数"
// @Security ApiKeyAuth
// @Success 200 {object} model.ResponseTmList
// @Router /admin/product/baseTrademark/{page}/{limit} [get]
func (*trademarkController) GetTrademark(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)

	// 获取数据
	data, err := service.TrademarkService.GetTrademarkList(page, size)
	if err != nil {
		zap.L().Error("service.TrademarkService.GetTrademarkList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

// UpdateTrademark 更新品牌
// @Summary 更新品牌接口
// @Description 处理更新品牌请求
// @Tags 品牌管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param object body model.ParamTmUpdate true "用户登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/product/baseTrademark/update [put]
func (*trademarkController) UpdateTrademark(c *gin.Context) {
	// 1. 获取参数及参数校验
	p := new(model.ParamTmUpdate)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("UpdateTrademark with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 业务处理
	err := service.TrademarkService.UpdateTrademark(p)
	if err != nil {
		zap.L().Error("service.TrademarkService.UpdateTrademark() failed", zap.String("tm_name", p.TmName), zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回结果
	ResponseSuccess(c, nil)
}

// DeleteTrademark 删除品牌
// @Summary 删除品牌接口
// @Description 处理删除品牌请求
// @Tags 品牌管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param id path int true "品牌 ID"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/product/baseTrademark/remove/{id} [delete]
func (*trademarkController) DeleteTrademark(c *gin.Context) {
	idStr := c.Param("id")
	tmId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("DeleteTrademark with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	err = service.TrademarkService.DeleteTrademark(tmId)
	if err != nil {
		zap.L().Error("service.TrademarkService.DeleteTrademark() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// GetAllTrademarkList 获取所有品牌列表
// @Summary 获取所有品牌列表
// @Description 获取所有品牌列表
// @Tags 品牌管理
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseAllTmList
// @Router /admin/product/baseTrademark/getTrademarkList [get]
func (*trademarkController) GetAllTrademarkList(c *gin.Context) {
	data, err := service.TrademarkService.GetAllTrademarkList()
	if err != nil {
		zap.L().Error("service.TrademarkService.GetAllTrademarkList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
