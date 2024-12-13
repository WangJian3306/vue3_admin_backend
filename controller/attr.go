package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"vue3_admin/model"
	"vue3_admin/service"
)

var AttrController attrController

type attrController struct {
}

// SaveAttrInfo 添加或者修改已有的属性
// @Summary 添加或者修改已有的属性的接口
// @Description 处理添加或者修改已有的属性请求
// @Tags 商品基础属性接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param object body model.Attr true "属性信息"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/product/saveAttrInfo [post]
func (*attrController) SaveAttrInfo(c *gin.Context) {
	p := new(model.Attr)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("CreateAttr with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 业务处理
	// 有属性ID则进行更新
	if p.AttrID != 0 {
		err := service.AttrService.UpdateAttr(p)
		if err != nil {
			zap.L().Error("service.AttrService.CreateAttr() failed", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
	} else {
		// 没有属性ID则进行创建
		err := service.AttrService.CreateAttr(p)
		if err != nil {
			zap.L().Error("service.AttrService.CreateAttr() failed", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
	}

	// 3. 返回结果
	ResponseSuccess(c, nil)
}

// GetAttr 获取分类下已有的属性与属性值
// @Summary 获取分类下已有的属性与属性值接口
// @Description 处理获取分类下已有的属性与属性值请求
// @Tags 商品基础属性接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param c1Id path int true "分类一 ID"
// @Param c2Id path int true "分类二 ID"
// @Param c3Id path int true "分类三 ID"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseToAssignMenuList
// @Router /admin/product/attrInfoList/{c1Id}/{c2Id}/{c3Id} [get]
func (*attrController) GetAttr(c *gin.Context) {
	c1Id, c2Id, c3Id, err := getAllCategoryId(c)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 获取数据
	data, err := service.AttrService.GetAttr(c1Id, c2Id, c3Id)
	if err != nil {
		zap.L().Error("service.AttrService.GetAttr failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

// DeleteAttr 删除基础属性
// @Summary 删除基础属性接口
// @Description 处理删除基础属性请求
// @Tags 商品基础属性接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param attrId path int true "属性 ID"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/deleteAttr/{attrId} [delete]
func (a *attrController) DeleteAttr(c *gin.Context) {
	idStr := c.Param("attrId")
	attrId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err = service.AttrService.DeleteAttr(attrId)
	if err != nil {
		zap.L().Error("service.AttrService.DeleteAttr() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}
