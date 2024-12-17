package service

import (
	"math"
	"vue3_admin/dao/mysql"
	"vue3_admin/model"
	"vue3_admin/pkg/snowflake"
)

var SkuService skuService

type skuService struct{}

func (s *skuService) SaveSkuInfo(skuInfo *model.SkuInfo) error {
	// SKU
	skuId := snowflake.GenID()
	spuId, err := skuInfo.SpuID.Int64()
	if err != nil {
		return err
	}
	c3Id, err := skuInfo.Category3ID.Int64()
	if err != nil {
		return err
	}
	tmId, err := skuInfo.TmID.Int64()
	if err != nil {
		return err
	}
	weight, err := skuInfo.Weight.Int64()
	if err != nil {
		return err
	}
	price, err := skuInfo.Price.Int64()
	if err != nil {
		return err
	}
	sku := &model.Sku{
		SkuID:         skuId,
		SpuID:         spuId,
		Category3ID:   c3Id,
		TmID:          tmId,
		SkuName:       skuInfo.SkuName,
		Weight:        weight,
		Price:         price,
		SkuDesc:       skuInfo.SkuDesc,
		SkuDefaultImg: skuInfo.SkuDefaultImg,
		IsSale:        skuInfo.IsSale,
	}

	// SKU 图片列表
	skuImageList := make([]*model.SkuImg, 0, len(skuInfo.SkuImageList))
	if len(skuInfo.SkuImageList) > 0 {
		for _, image := range skuInfo.SkuImageList {
			skuImageList = append(skuImageList, &model.SkuImg{
				ImageID:    snowflake.GenID(),
				SkuID:      skuId,
				ImageName:  image.ImageName,
				ImageURL:   image.ImageURL,
				SpuImageID: image.SpuImageID,
				IsDefault:  image.IsDefault,
			})
		}
	}

	// SKU 属性
	skuAttrValueList := make([]*model.SkuAttrValue, 0, len(skuInfo.SkuAttrValueList))
	if len(skuInfo.SkuAttrValueList) > 0 {
		for _, attrValue := range skuInfo.SkuAttrValueList {
			attrId, err := attrValue.AttrID.Int64()
			if err != nil {
				return err
			}
			valueId, err := attrValue.ValueID.Int64()
			if err != nil {
				return err
			}
			attrName, err := mysql.SkuDao.GetAttrNameByAttrId(attrId)
			if err != nil {
				return err
			}

			attrValueName, err := mysql.SkuDao.GetAttrValueNameByValueId(valueId)
			if err != nil {
				return err
			}
			skuAttrValueList = append(skuAttrValueList, &model.SkuAttrValue{
				SkuAttrValueID: snowflake.GenID(),
				AttrID:         attrId,
				ValueID:        valueId,
				ValueName:      attrValueName,
				AttrName:       attrName,
				SkuID:          skuId,
			})
		}
	}

	// SKU 销售属性
	skuSaleAttrValueList := make([]*model.SkuSaleAttrValue, 0, len(skuInfo.SkuSaleAttrValueList))
	if len(skuInfo.SkuSaleAttrValueList) > 0 {
		for _, saleAttrValue := range skuInfo.SkuSaleAttrValueList {
			saleAttrId, err := saleAttrValue.SaleAttrID.Int64()
			if err != nil {
				return err
			}
			saleAttrValueId, err := saleAttrValue.SaleAttrValueID.Int64()
			if err != nil {
				return err
			}
			saleAttrName, err := mysql.SkuDao.GetSaleAttrNameBySaleAttrId(saleAttrId)
			if err != nil {
				return err
			}
			saleAttrValueName, err := mysql.SkuDao.GetSaleAttrValueNameBySaleAttrValueId(saleAttrValueId)
			if err != nil {
				return err
			}

			skuSaleAttrValueList = append(skuSaleAttrValueList, &model.SkuSaleAttrValue{
				SkuSaleAttrValueID: snowflake.GenID(),
				SaleAttrID:         saleAttrId,
				SaleAttrName:       saleAttrName,
				SaleAttrValueID:    saleAttrValueId,
				SaleAttrValueName:  saleAttrValueName,
				SkuID:              skuId,
			})
		}
	}

	err = mysql.SkuDao.SaveSkuInfo(sku, skuImageList, skuAttrValueList, skuSaleAttrValueList)

	return err
}

func (s *skuService) FindBySpuId(spuId int64) ([]*model.ResponseSkuInfo, error) {
	return mysql.SkuDao.FindBySpuId(spuId)
}

func (s *skuService) GetSkuList(page, limit int64) (*model.ResponseSkuInfoList, error) {
	skuList, count, err := mysql.SkuDao.GetSkuList(page, limit)
	if err != nil {
		return nil, err
	}
	data := &model.ResponseSkuInfoList{
		Records:     skuList,
		Total:       count,
		Size:        limit,
		Current:     page,
		SearchCount: true,
		Pages:       int64(math.Ceil(float64(count) / float64(limit))),
	}
	return data, err
}

func (s *skuService) OnSaleSku(skuId int64) (err error) {
	return mysql.SkuDao.OnSaleSku(skuId)
}

func (s *skuService) CancelSaleSku(skuId int64) error {
	return mysql.SpuDao.CancelSaleSku(skuId)
}

func (s *skuService) DeleteSku(skuId int64) error {
	return mysql.SkuDao.DeleteSku(skuId)
}

func (s *skuService) GetSkuInfo(skuId int64) (skuInfo *model.ResponseSkuInfo, err error) {
	sku, err := mysql.SkuDao.GetSku(skuId)
	if err != nil {
		return nil, err
	}

	skuAttrValueList, err := mysql.SkuDao.GetSkuAttrValueList(skuId)
	if err != nil {
		return nil, err
	}
	skuSaleAttrValueList, err := mysql.SkuDao.GetSkuSaleAttrValueList(skuId)
	if err != nil {
		return nil, err
	}

	skuImageList, err := mysql.SkuDao.GetSkuImageList(skuId)

	skuInfo = &model.ResponseSkuInfo{
		Sku: model.Sku{
			SkuID:         sku.SkuID,
			SpuID:         sku.SpuID,
			Category3ID:   sku.Category3ID,
			TmID:          sku.SkuID,
			SkuName:       sku.SkuName,
			Weight:        sku.Weight,
			Price:         sku.Price,
			SkuDesc:       sku.SkuDesc,
			SkuDefaultImg: sku.SkuDefaultImg,
			IsSale:        sku.IsSale,
		},
		SkuAttrValueList:     skuAttrValueList,
		SkuSaleAttrValueList: skuSaleAttrValueList,
		SkuImageList:         skuImageList,
	}
	return skuInfo, err
}
