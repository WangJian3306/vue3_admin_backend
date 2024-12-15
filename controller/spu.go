package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"vue3_admin/model"
	"vue3_admin/service"
)

var SpuController spuController

type spuController struct{}

// GetSaleAttrList 获取所有销售列表
// @Summary 获取所有销售列表
// @Description 获取所有销售列表
// @Tags 商品SPU接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseAllSaleAttrList
// @Router /admin/product/baseSaleAttrList [get]
func (s *spuController) GetSaleAttrList(c *gin.Context) {
	saleAttrList, err := service.SpuService.GetSaleAttrList()
	if err != nil {
		zap.L().Error("service.SpuService.GetSaleAttrList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, saleAttrList)
}

// SaveSpuInfo 新增SPU接口
// @Summary 新增SPU接口
// @Description 新增SPU接口
// @Tags 商品SPU接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param object body model.Spu true "SPU信息"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/product/saveSpuInfo [post]
func (s *spuController) SaveSpuInfo(c *gin.Context) {
	p := new(model.Spu)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SaveSpu with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//fmt.Printf("%#v\n", p.SpuImageList)
	//for _, image := range p.SpuImageList {
	//	fmt.Println(image.ImageUrl)
	//}
	//for _, spuSaleAttrList := range p.SpuSaleAttrList {
	//	fmt.Printf("%#v\n", spuSaleAttrList)
	//	for _, spuSaleAttr := range spuSaleAttrList.SpuSaleAttrValue {
	//		fmt.Printf("%#v\n", spuSaleAttr)
	//	}
	//}

	err := service.SpuService.SaveSpuInfo(p)
	if err != nil {
		zap.L().Error("service.SpuService.SaveSpuInfo() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// GetSpuList 获取SPU分页列表（
// @Summary 获取SPU分页列表（
// @Description 获取SPU分页列表（
// @Tags 商品SPU接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param page path int true "当前页码"
// @Param limit path int true "每页记录数"
// @Param category3Id query int true "三级分类 ID"
// @Security ApiKeyAuth
// @Success 200 {object} model.ResponseSpuList
// @Router /admin/product/{page}/{limit} [get]
func (s *spuController) GetSpuList(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	// 获取三级分类id
	c3IdStr := c.Query("category3Id")
	c3Id, err := strconv.ParseInt(c3IdStr, 10, 64)
	if err != nil {
		zap.L().Error("service.SpuService.GetSpuList() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 获取列表数据
	responseSpuList, err := service.SpuService.GetSpuList(c3Id, page, size)
	if err != nil {
		zap.L().Error("service.SpuService.GetSpuList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, responseSpuList)

}

// GetSpuImageList 获取商品图片列表
// @Summary 获取商品图片列表接口
// @Description 处理获取商品图片列表请求
// @Tags 商品SKU接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param id path int true "SPU ID"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseSpuImageList
// @Router /admin/product/spuImageList/{id} [get]
func (s *spuController) GetSpuImageList(c *gin.Context) {
	idStr := c.Param("id")
	spuId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("service.SpuService.GetSpuImageList() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	spuImageList, err := service.SpuService.GetSpuImageList(spuId)
	if err != nil {
		zap.L().Error("service.SpuService.GetSpuImageList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, spuImageList)
}

// GetSpuSaleAttrList 获取商品销售属性列表
// @Summary 获取商品销售属性列表接口
// @Description 处理获取商品销售属性列表请求
// @Tags 商品SKU接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param id path int true "SPU ID"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseSpuSaleAttrList
// @Router /admin/product/spuSaleAttrList/{id} [get]
func (s *spuController) GetSpuSaleAttrList(c *gin.Context) {
	idStr := c.Param("id")
	spuId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("service.SpuService.GetSpuSaleAttrList() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	spuSaleAttrList, err := service.SpuService.GetSpuSaleAttrList(spuId)
	if err != nil {
		zap.L().Error("service.SpuService.GetSpuSaleAttrList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, spuSaleAttrList)
}

// UpdateSpuInfo 更新SPU接口
// @Summary 更新SPU接口
// @Description 更新SPU接口
// @Tags 商品SPU接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param object body model.Spu true "SPU信息"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/product/updateSpuInfo [post]
func (s *spuController) UpdateSpuInfo(c *gin.Context) {
	p := new(model.Spu)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateSpuInfo with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	err := service.SpuService.UpdateSpuInfo(p)
	if err != nil {
		zap.L().Error("service.SpuService.UpdateSpuInfo() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// DeleteSpu 删除SPU
// @Summary 删除SPU接口
// @Description 处理删除SPU请求
// @Tags 商品SPU接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户 Token"
// @Param spuId path int true "SPU ID"
// @Security ApiKeyAuth
// @Success 200 {object} ResponseData
// @Router /admin/product/deleteSpu/{spuId} [delete]
func (s *spuController) DeleteSpu(c *gin.Context) {
	idStr := c.Param("id")
	spuId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("service.SpuService.DeleteSpu() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	err = service.SpuService.DeleteSpu(spuId)
	if err != nil {
		zap.L().Error("service.SpuService.DeleteSpu() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
