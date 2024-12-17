package service

import (
	"math"
	"vue3_admin/dao/mysql"
	"vue3_admin/model"
	"vue3_admin/pkg/snowflake"
)

var SpuService spuService

type spuService struct {
}

func (s *spuService) GetSaleAttrList() ([]*model.SaleAttr, error) {
	return mysql.SpuDao.GetSaleAttrList()
}

func (s *spuService) SaveSpuInfo(p *model.Spu) (err error) {
	// SPU
	spuId := snowflake.GenID()
	spu := &model.Spu{
		SpuID:       spuId,
		SpuName:     p.SpuName,
		Description: p.Description,
		Category3ID: p.Category3ID,
		TmID:        p.TmID,
	}

	// 图片列表
	imageList := make([]*model.SpuImage, 0, 2)
	if len(p.SpuImageList) > 0 {
		for _, image := range p.SpuImageList {
			imageId := snowflake.GenID()
			imageList = append(imageList, &model.SpuImage{
				ImageID:   imageId,
				ImageName: image.ImageName,
				ImageUrl:  image.ImageUrl,
				SpuID:     spuId,
			})
		}
	}

	// SPU 销售属性
	spuSaleAttrList := make([]*model.SpuSaleAttr, 0, 2)
	if len(p.SpuSaleAttrList) > 0 {
		for _, spuSaleAttr := range p.SpuSaleAttrList {
			saleAttrId := snowflake.GenID()
			spuSaleAttrList = append(spuSaleAttrList, &model.SpuSaleAttr{
				SpuSaleAttrID:  saleAttrId,
				BaseSaleAttrId: spuSaleAttr.BaseSaleAttrId,
				SaleAttrName:   spuSaleAttr.SaleAttrName,
				SpuId:          spuId,
			})
		}
	}

	// SPU销售属性值
	spuSaleAttrValueList := make([]*model.SaleAttrValue, 0, 2)
	if len(p.SpuSaleAttrList) > 0 {
		for _, spuSaleAttr := range p.SpuSaleAttrList {
			for _, spuSaleAttrValue := range spuSaleAttr.SpuSaleAttrValue {
				saleAttrValueId := snowflake.GenID()
				spuSaleAttrValueList = append(spuSaleAttrValueList, &model.SaleAttrValue{
					SaleAttrValueID:   saleAttrValueId,
					SaleAttrValueName: spuSaleAttrValue.SaleAttrValueName,
					BaseSaleAttrId:    spuSaleAttrValue.BaseSaleAttrId,
					SpuId:             spuId,
				})
			}
		}
	}

	err = mysql.SpuDao.SaveSpuInfo(spu, imageList, spuSaleAttrList, spuSaleAttrValueList)
	return err
}

func (s *spuService) GetSpuList(c3Id, page, limit int64) (spu *model.ResponseSpuList, err error) {
	spuList, count, err := mysql.SpuDao.GetSpuList(c3Id, page, limit)
	if err != nil {
		return nil, err
	}

	spu = &model.ResponseSpuList{
		Records:     spuList,
		Total:       count,
		Size:        limit,
		Current:     page,
		SearchCount: true,
		Pages:       int64(math.Ceil(float64(count) / float64(limit))),
	}

	return spu, err
}

func (s *spuService) GetSpuImageList(spuId int64) (spuImageList []*model.SpuImage, err error) {
	return mysql.SpuDao.GetSpuImageList(spuId)
}

func (s *spuService) GetSpuSaleAttrList(spuId int64) (spuSaleAttrList []*model.SpuSaleAttr, err error) {
	spuSaleAttrList, err = mysql.SpuDao.GetSpuSaleAttrList(spuId)
	if spuSaleAttrList != nil && len(spuSaleAttrList) > 0 {
		for _, spuSaleAttr := range spuSaleAttrList {
			saleAttrValueList, err := mysql.SpuDao.GetSpuSaleAttrValueList(spuId, spuSaleAttr.BaseSaleAttrId)
			if err != nil {
				return nil, err
			}
			spuSaleAttr.SpuSaleAttrValue = saleAttrValueList
		}
	}
	return spuSaleAttrList, err
}

func (s *spuService) UpdateSpuInfo(spu *model.Spu) error {
	// 创建一个 SPU 对象
	newSpu := &model.Spu{
		SpuID:       spu.SpuID,
		SpuName:     spu.SpuName,
		Description: spu.Description,
		Category3ID: spu.Category3ID,
		TmID:        spu.TmID,
	}

	// 创建一个图片列表
	imageList := make([]*model.SpuImage, 0, 2)
	if len(spu.SpuImageList) > 0 {
		for _, image := range spu.SpuImageList {
			imageId := snowflake.GenID()
			imageList = append(imageList, &model.SpuImage{
				ImageID:   imageId,
				ImageName: image.ImageName,
				ImageUrl:  image.ImageUrl,
				SpuID:     spu.SpuID,
			})
		}
	}

	// SPU 销售属性
	spuSaleAttrList := make([]*model.SpuSaleAttr, 0, 2)
	if len(spu.SpuSaleAttrList) > 0 {
		for _, spuSaleAttr := range spu.SpuSaleAttrList {
			spuSaleAttrId := snowflake.GenID()
			spuSaleAttrList = append(spuSaleAttrList, &model.SpuSaleAttr{
				SpuSaleAttrID:  spuSaleAttrId,
				BaseSaleAttrId: spuSaleAttr.BaseSaleAttrId,
				SaleAttrName:   spuSaleAttr.SaleAttrName,
				SpuId:          spu.SpuID,
			})
		}
	}

	// SPU销售属性值
	spuSaleAttrValueList := make([]*model.SaleAttrValue, 0, 2)
	if len(spu.SpuSaleAttrList) > 0 {
		for _, spuSaleAttr := range spu.SpuSaleAttrList {
			if spuSaleAttr.SpuSaleAttrValue != nil && len(spuSaleAttr.SpuSaleAttrValue) > 0 {
				for _, spuSaleAttrValue := range spuSaleAttr.SpuSaleAttrValue {
					saleAttrValueId := snowflake.GenID()
					spuSaleAttrValueList = append(spuSaleAttrValueList, &model.SaleAttrValue{
						SaleAttrValueID:   saleAttrValueId,
						SaleAttrValueName: spuSaleAttrValue.SaleAttrValueName,
						BaseSaleAttrId:    spuSaleAttrValue.BaseSaleAttrId,
						SpuId:             spu.SpuID,
					})
				}
			}

		}
	}

	err := mysql.SpuDao.UpdateSpuInfo(newSpu, imageList, spuSaleAttrList, spuSaleAttrValueList)
	if err != nil {
		return err
	}
	return nil
}

func (s *spuService) DeleteSpu(spuId int64) error {
	return mysql.SpuDao.DeleteSpu(spuId)
}
