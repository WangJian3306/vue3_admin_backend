package model

type SaleAttr struct {
	BaseModel

	SaleAttrID   int64  `json:"id" db:"sale_attr_id"`     // 销售属性ID
	SaleAttrName string `json:"name" db:"sale_attr_name"` // 销售属性名称
}

type SaleAttrValue struct {
	BaseModel
	SaleAttrValueID   int64  `json:"id" db:"sale_attr_value_id"`
	SaleAttrValueName string `json:"saleAttrValueName" db:"sale_attr_value_name"`
	BaseSaleAttrId    int64  `json:"baseSaleAttrId" db:"sale_attr_id"`
	SpuId             int64  `json:"spuId" db:"spu_id"`
}

type SpuSaleAttr struct {
	SpuSaleAttrID    int64            `json:"id" db:"spu_sale_attr_id"`
	BaseSaleAttrId   int64            `json:"baseSaleAttrId" db:"base_sale_attr_id"`
	SaleAttrName     string           `json:"saleAttrName" db:"sale_attr_name"`
	SpuId            int64            `json:"spuId" db:"spu_id"`
	SpuSaleAttrValue []*SaleAttrValue `json:"spuSaleAttrValueList"`
}
type SpuImage struct {
	BaseModel
	ImageID   int64  `json:"id" db:"image_id"`
	ImageName string `json:"imgName" db:"image_name"`
	ImageUrl  string `json:"imgUrl" db:"image_url"`
	SpuID     int64  `json:"spuId" db:"spu_id"`
}

type Spu struct {
	BaseModel
	SpuID           int64          `json:"id" db:"spu_id"`
	SpuName         string         `json:"spuName" db:"spu_name"`
	Description     string         `json:"description" db:"description"`
	Category3ID     int64          `json:"category3Id" db:"category3_id"`
	TmID            int64          `json:"tmId" db:"tm_id"`
	SpuImageList    []*SpuImage    `json:"spuImageList" db:"-"`
	SpuSaleAttrList []*SpuSaleAttr `json:"spuSaleAttrList" db:"-"`
}

type ResponseSpuList struct {
	Records     []*Spu `json:"records"`
	Total       int64  `json:"total"`
	Size        int64  `json:"size"`
	Current     int64  `json:"current"`
	SearchCount bool   `json:"searchCount"`
	Pages       int64  `json:"pages"`
}
