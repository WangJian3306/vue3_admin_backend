package model

import "encoding/json"

type Sku struct {
	BaseModel
	SkuID         int64  `json:"id" db:"sku_id"`
	SpuID         int64  `json:"spuID" db:"spu_id"`
	Category3ID   int64  `json:"category3Id" db:"category_3_id"`
	TmID          int64  `json:"tmId" db:"tm_id"`
	SkuName       string `json:"skuName" db:"sku_name"`
	Weight        int64  `json:"weight" db:"weight"`
	Price         int64  `json:"price" db:"price"`
	SkuDesc       string `json:"skuDesc" db:"sku_desc"`
	SkuDefaultImg string `json:"skuDefaultImg" db:"sku_default_img"`
	IsSale        int8   `json:"isSale" db:"is_sale"`
}

type SkuAttrValue struct {
	BaseModel
	SkuAttrValueID int64  `json:"id" db:"sku_attr_value_id"`
	AttrID         int64  `json:"attrId" db:"attr_id"`   // 平台属性ID
	ValueID        int64  `json:"valueId" db:"value_id"` // 属性值ID
	ValueName      string `json:"valueName" db:"value_name"`
	AttrName       string `json:"attrName" db:"attr_name"`
	SkuID          int64  `json:"skuId" db:"sku_id"`
}

type SkuSaleAttrValue struct {
	BaseModel
	SkuSaleAttrValueID int64  `json:"id" db:"sku_sale_attr_value_id"`
	SaleAttrID         int64  `json:"saleAttrId" db:"sale_attr_id"` // 销售属性ID
	SaleAttrValueID    int64  `json:"saleAttrValueId" db:"sale_attr_value_id"`
	SaleAttrName       string `json:"saleAttrName" db:"sale_attr_name"`
	SaleAttrValueName  string `json:"saleAttrValueName" db:"sale_attr_value_name"`
	SkuID              int64  `json:"skuId" db:"sku_id"`
}

type SkuImg struct {
	BaseModel
	ImageID    int64  `json:"id" db:"image_id"`
	SkuID      int64  `json:"skuId" db:"sku_id"`
	ImageName  string `json:"imgName" db:"image_name"`
	ImageURL   string `json:"imgUrl" db:"image_url"`
	SpuImageID int64  `json:"spuImgId" db:"spu_image_id"`
	IsDefault  string `json:"isDefault" db:"is_default"`
}

type SkuAttrValueDTO struct {
	SkuAttrValueID int64       `json:"id"`
	AttrID         json.Number `json:"attrId"`  // 平台属性ID
	ValueID        json.Number `json:"valueId"` // 属性值ID
	ValueName      string      `json:"valueName"`
	AttrName       string      `json:"attrName"`
	SkuID          int64       `json:"skuId"`
}

type SkuSaleAttrValueDTO struct {
	SkuSaleAttrValueID int64       `json:"id"`
	SaleAttrID         json.Number `json:"saleAttrId"` // 销售属性ID
	SaleAttrValueID    json.Number `json:"saleAttrValueId"`
	SaleAttrName       string      `json:"saleAttrName"`
	SaleAttrValueName  string      `json:"saleAttrValueName"`
	SkuID              int64       `json:"skuId"`
}

type SkuImgDTO struct {
	ImageID    int64  `json:"id"`
	SkuID      int64  `json:"skuId"`
	ImageName  string `json:"imgName"`
	ImageURL   string `json:"imgUrl"`
	SpuImageID int64  `json:"spuImgId"`
	IsDefault  string `json:"isDefault"`
}

type SkuInfo struct {
	SkuID                int64                  `json:"id"`
	SpuID                json.Number            `json:"spuID"`
	Category3ID          json.Number            `json:"category3Id"`
	TmID                 json.Number            `json:"tmId"`
	SkuName              string                 `json:"skuName"`
	Weight               json.Number            `json:"weight"`
	Price                json.Number            `json:"price"`
	SkuDesc              string                 `json:"skuDesc"`
	SkuDefaultImg        string                 `json:"skuDefaultImg"`
	IsSale               int8                   `json:"isSale"`
	SkuAttrValueList     []*SkuAttrValueDTO     `json:"skuAttrValueList"`
	SkuSaleAttrValueList []*SkuSaleAttrValueDTO `json:"skuSaleAttrValueList"`
	SkuImageList         []*SkuImgDTO           `json:"skuImageList"`
}

type ResponseSkuInfo struct {
	Sku
	SkuAttrValueList     []*SkuAttrValue     `json:"skuAttrValueList"`
	SkuSaleAttrValueList []*SkuSaleAttrValue `json:"skuSaleAttrValueList"`
	SkuImageList         []*SkuImg           `json:"skuImageList"`
}

type ResponseSkuInfoList struct {
	Records     []*ResponseSkuInfo `json:"records"`
	Total       int64              `json:"total"`
	Size        int64              `json:"size"`
	Current     int64              `json:"current"`
	SearchCount bool               `json:"searchCount"`
	Pages       int64              `json:"pages"`
}
