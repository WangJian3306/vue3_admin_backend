package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"vue3_admin/model"
)

var SkuDao skuDao

type skuDao struct{}

func (*skuDao) SaveSkuInfo(sku *model.Sku, skuImageList []*model.SkuImg, skuAttrValueList []*model.SkuAttrValue, skuSaleAttrValueList []*model.SkuSaleAttrValue) error {
	// 开启事务
	tx, err := db.Beginx()
	if err != nil {
		zap.L().Error("begin trans failed", zap.Error(err))
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()

	// 保存 SKU
	skuSqlStr := `INSERT INTO sku (sku_id, spu_id, category_3_id, tm_id, sku_name, weight, price, sku_desc, sku_default_img, 
                 is_sale) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	rs, err := tx.Exec(skuSqlStr, sku.SkuID, sku.SpuID, sku.Category3ID, sku.TmID, sku.SkuName, sku.Weight, sku.Price, sku.SkuDesc, sku.SkuDefaultImg, sku.IsSale)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 保存 SKU 图片
	if skuImageList != nil && len(skuImageList) > 0 {
		rs, err = tx.NamedExec("INSERT INTO sku_image (image_id, sku_id, image_url, spu_image_id, is_default) VALUES (:image_id, :sku_id, :image_url, :spu_image_id, :is_default)", skuImageList)
		if err != nil {
			return err
		}
		_, err = rs.RowsAffected()
		if err != nil {
			return err
		}
	}

	// 保存 SKU 属性
	if skuAttrValueList != nil && len(skuAttrValueList) > 0 {
		rs, err = tx.NamedExec("INSERT INTO sku_attr_value (sku_attr_value_id, attr_id, value_id, value_name, attr_name, sku_id) VALUES (:sku_attr_value_id, :attr_id, :value_id, :value_name, :attr_name, :sku_id)", skuAttrValueList)
		if err != nil {
			return err
		}
		_, err = rs.RowsAffected()
		if err != nil {
			return err
		}
	}

	// 保存 SKU 销售属性
	if skuSaleAttrValueList != nil && len(skuSaleAttrValueList) > 0 {
		rs, err = tx.NamedExec("INSERT INTO sku_sale_attr_value (sku_sale_attr_value_id, sale_attr_id, sale_attr_value_id, sale_attr_name, sale_attr_value_name, sku_id) VALUES (:sku_sale_attr_value_id, :sale_attr_id, :sale_attr_value_id, :sale_attr_name, :sale_attr_value_name, :sku_id)", skuSaleAttrValueList)
		if err != nil {
			return err
		}
		_, err = rs.RowsAffected()
		if err != nil {
			return err
		}
	}

	return nil
}

func (*skuDao) GetSkuList(page, limit int64) ([]*model.ResponseSkuInfo, int64, error) {
	countSqlStr := `SELECT count(sku_id) from sku`
	var count int64
	if err := db.Get(&count, countSqlStr); err != nil {
		return nil, 0, err
	}
	sqlStr := `SELECT sku_id, spu_id, category_3_id, tm_id, sku_name, weight, price, sku_desc, sku_default_img, is_sale FROM sku LIMIT ?,?`
	skuList := make([]*model.ResponseSkuInfo, 0, limit)
	if err := db.Select(&skuList, sqlStr, (page-1)*limit, limit); err != nil {
		return nil, 0, err
	}
	return skuList, count, nil
}

func (*skuDao) OnSaleSku(skuId int64) (err error) {
	sqlStr := `UPDATE sku SET is_sale = true WHERE sku_id = ?`
	_, err = db.Exec(sqlStr, skuId)
	return err
}

func (s *skuDao) FindBySpuId(spuId int64) ([]*model.ResponseSkuInfo, error) {
	sqlStr := `SELECT sku_id, spu_id, category_3_id, tm_id, sku_name, weight, price, sku_desc, sku_default_img, is_sale FROM sku WHERE spu_id = ?`
	skuList := make([]*model.ResponseSkuInfo, 0)
	if err := db.Select(&skuList, sqlStr, spuId); err != nil {
		return nil, err
	}
	return skuList, nil
}

func (s *spuDao) CancelSaleSku(skuId int64) (err error) {
	sqlStr := `UPDATE sku SET is_sale = false WHERE sku_id = ?`
	_, err = db.Exec(sqlStr, skuId)
	return err
}

func (s *skuDao) DeleteSku(skuId int64) (err error) {
	// 开启事务
	tx, err := db.Beginx()
	if err != nil {
		zap.L().Error("begin trans failed", zap.Error(err))
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()

	// 删除 SKU
	rs, err := tx.Exec("DELETE FROM sku WHERE sku_id = ?", skuId)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 删除图片列表
	rs, err = tx.Exec("DELETE FROM sku_image WHERE sku_id = ?", skuId)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 删除属性列表
	rs, err = tx.Exec("DELETE FROM sku_attr_value WHERE sku_id = ?", skuId)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 删除消暑属性列表
	rs, err = tx.Exec("DELETE FROM sku_sale_attr_value WHERE sku_id = ?", skuId)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *skuDao) GetSku(skuId int64) (sku *model.Sku, err error) {
	sqlStr := `SELECT id, sku_id, spu_id, category_3_id, tm_id, sku_name, weight, price, sku_desc, sku_default_img, is_sale FROM sku WHERE sku_id = ?`
	sku = new(model.Sku)
	if err := db.Get(sku, sqlStr, skuId); err != nil {
		return nil, err
	}
	return sku, nil
}

func (s *skuDao) GetSkuAttrValueList(skuId int64) (skuAttrValueList []*model.SkuAttrValue, err error) {
	sqlStr := `SELECT id, sku_attr_value_id, attr_id, value_id, value_name, attr_name, sku_id FROM sku_attr_value WHERE sku_id = ?`
	skuAttrValueList = make([]*model.SkuAttrValue, 0)
	if err := db.Select(&skuAttrValueList, sqlStr, skuId); err != nil {
		return nil, err
	}
	return skuAttrValueList, nil
}

func (s *skuDao) GetSkuSaleAttrValueList(skuId int64) (skuSaleAttrValueList []*model.SkuSaleAttrValue, err error) {
	sqlStr := `SELECT id, sku_sale_attr_value_id, sale_attr_id, sale_attr_value_id, sale_attr_name, sale_attr_value_name, sku_id FROM sku_sale_attr_value WHERE sku_id = ?`
	skuSaleAttrValueList = make([]*model.SkuSaleAttrValue, 0)
	if err := db.Select(&skuSaleAttrValueList, sqlStr, skuId); err != nil {
		return nil, err
	}
	return skuSaleAttrValueList, nil
}

func (s *skuDao) GetSkuImageList(skuId int64) (skuImageList []*model.SkuImg, err error) {
	sqlStr := `SELECT id, image_id, sku_id, image_url, spu_image_id, is_default FROM sku_image WHERE sku_id = ?`
	skuImageList = make([]*model.SkuImg, 0)
	if err := db.Select(&skuImageList, sqlStr, skuId); err != nil {
		return nil, err
	}
	return skuImageList, nil
}

func (s *skuDao) GetAttrNameByAttrId(attrId int64) (attrName string, err error) {
	sqlStr := fmt.Sprintf(`SELECT attr_name FROM attr WHERE attr_id = ?`)
	if err = db.Get(&attrName, sqlStr, attrId); err != nil {
		zap.L().Error("GetAttrNameByAttrId", zap.Error(err))
		return "", err
	}
	return attrName, nil
}

func (s *skuDao) GetAttrValueNameByValueId(valueId int64) (attrValueName string, err error) {
	sqlStr := fmt.Sprintf(`SELECT value_name FROM attr_value WHERE attr_value_id = ?`)
	if err = db.Get(&attrValueName, sqlStr, valueId); err != nil {
		zap.L().Error("GetAttrValueNameByValueId", zap.Error(err))
		return "", err
	}
	return attrValueName, nil
}

func (s *skuDao) GetSaleAttrNameBySaleAttrId(saleAttrId int64) (saleAttrName string, err error) {
	sqlStr := `SELECT sale_attr_name FROM spu_sale_attr WHERE spu_sale_attr_id = ?`
	if err = db.Get(&saleAttrName, sqlStr, saleAttrId); err != nil {
		zap.L().Error("GetSaleAttrNameBySaleAttrId", zap.Error(err))
		return "", err
	}
	return saleAttrName, nil
}

func (s *skuDao) GetSaleAttrValueNameBySaleAttrValueId(saleAttrValueId int64) (saleAttrValueName string, err error) {
	sqlStr := `SELECT sale_attr_value_name FROM sale_attr_value WHERE sale_attr_value_id = ?`
	if err = db.Get(&saleAttrValueName, sqlStr, saleAttrValueId); err != nil {
		zap.L().Error("GetSaleAttrValueNameBySaleAttrValueId", zap.Error(err))
		return "", err
	}
	return saleAttrValueName, nil
}
