package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"vue3_admin/model"
)

var SpuDao spuDao

type spuDao struct{}

func (s *spuDao) GetSaleAttrList() ([]*model.SaleAttr, error) {
	sqlStr := `SELECT sale_attr_id, sale_attr_name FROM sale_attr`
	data := make([]*model.SaleAttr, 0, 3)
	if err := db.Select(&data, sqlStr); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *spuDao) SaveSpuInfo(spu *model.Spu, imageList []*model.SpuImage, spuSaleAttrList []*model.SpuSaleAttr, spuSaleAttrValueList []*model.SaleAttrValue) error {
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

	// 1. 保存 SPU
	spuSqlStr := `INSERT INTO spu (spu_id, spu_name, description, category3_id, tm_id) VALUES (?, ?, ?, ?, ?)`
	rs, err := tx.Exec(spuSqlStr, spu.SpuID, spu.SpuName, spu.Description, spu.Category3ID, spu.TmID)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 2. 保存图片列表
	rs, err = tx.NamedExec("INSERT INTO spu_image_list (image_id, image_name, image_url, spu_id) VALUES (:image_id, :image_name, :image_url, :spu_id)",
		imageList)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 3. 保存SPU销售属性
	rs, err = tx.NamedExec("INSERT INTO spu_sale_attr (spu_sale_attr_id, base_sale_attr_id, sale_attr_name, spu_id) VALUES (:spu_sale_attr_id, :base_sale_attr_id, :sale_attr_name, :spu_id)", spuSaleAttrList)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 4. 保存SPU销售属性值
	rs, err = tx.NamedExec("INSERT INTO sale_attr_value (sale_attr_value_id, sale_attr_value_name, sale_attr_id, spu_id) VALUES (:sale_attr_value_id, :sale_attr_value_name, :sale_attr_id, :spu_id)",
		spuSaleAttrValueList)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *spuDao) GetSpuList(c3Id, page, limit int64) ([]*model.Spu, int64, error) {
	countSqlStr := `SELECT count(spu_id) from spu`
	var count int64
	if err := db.Get(&count, countSqlStr); err != nil {
		return nil, 0, err
	}
	sqlStr := `SELECT spu_id, spu_name, description, category3_id, tm_id FROM spu WHERE category3_id = ? LIMIT ?,?`
	spuList := make([]*model.Spu, 0, limit)
	if err := db.Select(&spuList, sqlStr, c3Id, (page-1)*limit, limit); err != nil {
		return nil, 0, err
	}
	return spuList, count, nil
}

func (s *spuDao) GetSpuImageList(spuId int64) (spuImageList []*model.SpuImage, err error) {
	sqlStr := `SELECT image_id, image_name, image_url, spu_id FROM spu_image_list WHERE spu_id = ?`
	err = db.Select(&spuImageList, sqlStr, spuId)
	if err != nil {
		return nil, err
	}
	return spuImageList, nil
}

func (s *spuDao) GetSpuSaleAttrList(spuId int64) (saleAttrList []*model.SpuSaleAttr, err error) {
	sqlStr := `SELECT spu_sale_attr_id, base_sale_attr_id, sale_attr_name, spu_id FROM spu_sale_attr WHERE spu_id = ?`
	err = db.Select(&saleAttrList, sqlStr, spuId)
	if err != nil {
		return nil, err
	}
	return saleAttrList, nil
}

func (s *spuDao) GetSpuSaleAttrValueList(spuId, baseSaleAttrId int64) (saleAttrValueList []*model.SaleAttrValue, err error) {
	sqlStr := `SELECT sale_attr_value_id, sale_attr_value_name, sale_attr_id, spu_id FROM sale_attr_value WHERE spu_id = ? AND sale_attr_id = ?`
	err = db.Select(&saleAttrValueList, sqlStr, spuId, baseSaleAttrId)
	if err != nil {
		return nil, err
	}
	return saleAttrValueList, nil
}

func (s *spuDao) UpdateSpuInfo(spu *model.Spu, imageList []*model.SpuImage, spuSaleAttrList []*model.SpuSaleAttr, spuSaleAttrValueList []*model.SaleAttrValue) error {
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

	// 1. 更新 SPU
	spuSqlStr := `UPDATE spu SET spu_name = ?, description = ?, category3_id = ?, tm_id = ? WHERE spu_id = ?`
	rs, err := tx.Exec(spuSqlStr, spu.SpuName, spu.Description, spu.Category3ID, spu.TmID, spu.SpuID)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 2. 跟新图片列表
	// 2.1 先删除所有图片
	rs, err = tx.Exec("DELETE FROM spu_image_list WHERE spu_id = ?", spu.SpuID)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	// 2.2 插入图片列表
	rs, err = tx.NamedExec("INSERT INTO spu_image_list (image_id, image_name, image_url, spu_id) VALUES (:image_id, :image_name, :image_url, :spu_id)",
		imageList)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 3. 更新SPU销售属性
	// 3.1 删除所有销售属性
	rs, err = tx.Exec("DELETE FROM spu_sale_attr WHERE spu_id = ?", spu.SpuID)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	// 3.2 保存所有销售属性
	rs, err = tx.NamedExec("INSERT INTO spu_sale_attr (spu_sale_attr_id, base_sale_attr_id, sale_attr_name, spu_id) VALUES (:spu_sale_attr_id, :base_sale_attr_id, :sale_attr_name, :spu_id)", spuSaleAttrList)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 4. 更新SPU销售属性值
	// 4.1 删除所有SPU销售属性值
	rs, err = tx.Exec("DELETE FROM sale_attr_value WHERE spu_id = ?", spu.SpuID)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}
	// 4.2  保存所有SPU销售属性值
	rs, err = tx.NamedExec("INSERT INTO sale_attr_value (sale_attr_value_id, sale_attr_value_name, sale_attr_id, spu_id) VALUES (:sale_attr_value_id, :sale_attr_value_name, :sale_attr_id, :spu_id)",
		spuSaleAttrValueList)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *spuDao) DeleteSpu(spuId int64) error {
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

	// 删除 SPU
	rs, err := tx.Exec("DELETE FROM spu WHERE spu_id = ?", spuId)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 删除图片列表
	rs, err = tx.Exec("DELETE FROM spu_image_list WHERE spu_id = ?", spuId)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 删除属性列表
	rs, err = tx.Exec("DELETE FROM spu_sale_attr WHERE spu_id = ?", spuId)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 删除属性值列表
	rs, err = tx.Exec("DELETE FROM sale_attr_value WHERE spu_id = ?", spuId)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	return nil

}
