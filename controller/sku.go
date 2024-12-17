package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"vue3_admin/model"
	"vue3_admin/service"
)

var SkuController skuController

type skuController struct{}

// SaveSkuInfo 新增SKU接口
// @Summary 新增SKU接口
// @Description 新增SKU接口
// @Tags 商品SKU接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param object body model.SkuInfo true "SKU信息"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/product/saveSkuInfo [post]
func (s *skuController) SaveSkuInfo(c *gin.Context) {
	p := new(model.SkuInfo)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SaveSku with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//fmt.Printf("%#v\n", p)
	//fmt.Printf("%T , %v\n", p.Price, p.Price)
	//for _, value := range p.SkuAttrValueList {
	//	fmt.Println(value.AttrID)
	//}
	err := service.SkuService.SaveSkuInfo(p)
	if err != nil {
		zap.L().Error("service.SkuService.SaveSkuInfo() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// FindBySpuId 根据 SPU ID 查询 SKU 列表
// @Summary 根据 SPU ID 查询 SKU 接口
// @Description 处理根据 SPU ID 查询 SKU列表请求
// @Tags 商品SKU接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param id path int true "SPU ID"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseSkuFindBySpuId
// @Router /admin/product/findBySpuId/{spuId} [get]
func (s *skuController) FindBySpuId(c *gin.Context) {
	idStr := c.Param("id")
	spuId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("service.SkuController.FindBySpuId() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	skuList, err := service.SkuService.FindBySpuId(spuId)
	if err != nil {
		zap.L().Error("service.SkuController.FindBySpuId() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, skuList)
}

// GetSkuList 获取SKU分页列表（
// @Summary 获取SKU分页列表（
// @Description 获取SKU分页列表（
// @Tags 商品SKU接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param page path int true "当前页码"
// @Param limit path int true "每页记录数"
// @Security ApiKeyAuth
// @Success 200 {object} model.ResponseSkuInfoList
// @Router /admin/product/list/{page}/{limit} [get]
func (s *skuController) GetSkuList(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)

	skuList, err := service.SkuService.GetSkuList(page, size)

	if err != nil {
		zap.L().Error("service.SkuController.GetSkuList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, skuList)
}

// OnSaleSku 上架SKU
// @Summary 上架SKU接口
// @Description 处理上架SKU请求
// @Tags 商品SKU接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param skuId path int true "SKU ID"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/product/onSale/{skuId} [get]
func (s *skuController) OnSaleSku(c *gin.Context) {
	idStr := c.Param("id")
	skuId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("service.SkuController.OnSaleSku() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err = service.SkuService.OnSaleSku(skuId)
	if err != nil {
		zap.L().Error("service.SkuController.OnSaleSku() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}

// CancelSaleSku 下架SKU
// @Summary 下架SKU接口
// @Description 处理下架SKU请求
// @Tags 商品SKU接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param skuId path int true "SKU ID"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/product/cancelSale/{skuId} [get]
func (s *skuController) CancelSaleSku(c *gin.Context) {
	idStr := c.Param("id")
	skuId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("service.SkuController.CancelSaleSku() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err = service.SkuService.CancelSaleSku(skuId)
	if err != nil {
		zap.L().Error("service.SkuController.CancelSaleSku() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}

// DeleteSku 删除SKU
// @Summary 删除SKU接口
// @Description 处理删除SKU请求
// @Tags 商品SKU接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param skuId path int true "SKU ID"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/product/deleteSku/{skuId} [delete]
func (s *skuController) DeleteSku(c *gin.Context) {
	idStr := c.Param("id")
	skuId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("service.SkuController.DeleteSku() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err = service.SkuService.DeleteSku(skuId)
	if err != nil {
		zap.L().Error("service.SkuController.DeleteSku() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func (s *skuController) GetSkuInfo(c *gin.Context) {
	idStr := c.Param("id")
	skuId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("service.SkuController.GetSkuInfo() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	skuInfo, err := service.SkuService.GetSkuInfo(skuId)
	if err != nil {
		zap.L().Error("service.SkuController.GetSkuInfo() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, skuInfo)
}
