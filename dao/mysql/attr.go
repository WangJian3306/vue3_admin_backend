package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"vue3_admin/model"
	"vue3_admin/pkg/snowflake"
)

var AttrDao attrDao

type attrDao struct {
}

//func (a *attrDao) InsertAttr(attr *model.Attr) (err error) {
//	sqlStr := `INSERT INTO attr (attr_id, attr_name, category_id, category_level) VALUES (?,?,?,?)`
//	_, err = db.Exec(sqlStr, attr.AttrID, attr.AttrName, attr.CategoryId, attr.CategoryLevel)
//	return err
//}
//
//func (a *attrDao) InsertAttrValue(attrValues []*model.AttrValue) (err error) {
//
//	_, err = db.NamedExec(
//		"INSERT INTO attr_value (attr_value_id, value_name, attr_id) VALUES (:attr_value_id,:value_name,:attr_id)",
//		attrValues,
//	)
//
//	return err
//}

func (a *attrDao) UpdateAttrAndAttrValue(attr *model.Attr) (err error) {

	tx, err := db.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
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

	// 更新 Attr
	sqlStr1 := `UPDATE attr SET attr_name = ? WHERE attr_id = ?`
	rs, err := tx.Exec(sqlStr1, attr.AttrName, attr.AttrID)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 更新 AttrValue
	var attrValueIDs []int64
	for _, attrValue := range attr.AttrValueList {
		if attrValue.AttrValueID != 0 {
			attrValueIDs = append(attrValueIDs, attrValue.AttrValueID)
			_, err = tx.NamedExec(
				"UPDATE attr_value SET value_name = :value_name WHERE attr_value_id = :attr_value_id",
				attrValue,
			)
			if err != nil {
				return err
			}
			_, err = rs.RowsAffected()
			if err != nil {
				return err
			}
		} else {
			attrValue.AttrValueID = snowflake.GenID()
			attrValue.AttrID = attr.AttrID
			attrValueIDs = append(attrValueIDs, attrValue.AttrValueID)
			// 插入 AttrValue
			_, err = tx.NamedExec(
				"INSERT INTO attr_value (attr_value_id, value_name, attr_id) VALUES (:attr_value_id,:value_name,:attr_id)",
				attrValue,
			)
			if err != nil {
				return err
			}
			_, err = rs.RowsAffected()
			if err != nil {
				return err
			}
		}
	}

	// 删除 AttrValue
	query, args, err := sqlx.In("DELETE FROM attr_value WHERE attr_id in (?) AND attr_value_id NOT IN (?)", attr.AttrID, attrValueIDs)
	if err != nil {
		return
	}
	query = tx.Rebind(query)
	ret, err := tx.Exec(query, args...)
	if err != nil {
		zap.L().Error("delete failed", zap.Error(err))
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		zap.L().Error("get RowsAffected failed", zap.Error(err))
		return
	}
	zap.L().Info("delete success", zap.Int64("affected rows", n))

	return err
}

func (a *attrDao) InsertAttrAndAttrValue(attr *model.Attr) (err error) {

	tx, err := db.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
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

	sqlStr1 := `INSERT INTO attr (attr_id, attr_name, category_id, category_level) VALUES (?,?,?,?)`
	rs, err := tx.Exec(sqlStr1, attr.AttrID, attr.AttrName, attr.CategoryId, attr.CategoryLevel)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	rs, err = tx.NamedExec(
		"INSERT INTO attr_value (attr_value_id, value_name, attr_id) VALUES (:attr_value_id,:value_name,:attr_id)",
		attr.AttrValueList,
	)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	return err
}

func (a *attrDao) GetAttr(c1Id, c2Id, c3Id int64) (data []*model.Attr, err error) {
	//sqlStr := `SELECT attr_id, attr_name, category_id, category_level FROM attr WHERE category_id = ?`

	//var data = new(model.Attr)
	// 这里SQL不用写这么复杂，直接根据 c3Id 查 attr 表即可，这里写复杂纯粹是为了想把 c1Id 和 c2Id 用上
	sqlStr := `SELECT attr_id, attr_name, category_id, category_level FROM attr c
	    JOIN category3 a on c.category_id = a.category3_id
	    JOIN category2 b ON a.category2_id = b.category2_id AND b.category1_id = ? AND a.category2_id = ? AND a.category3_id = ?;`

	err = db.Select(&data, sqlStr, c1Id, c2Id, c3Id)
	if err != nil {
		return data, err
	}

	return data, err
}

func (a *attrDao) GetAttrValue(attrId int64) (data []*model.AttrValue, err error) {
	data = make([]*model.AttrValue, 0)
	sqlStr := `SELECT attr_value_id, value_name, attr_id FROM attr_value WHERE attr_id = ?`

	if err = db.Select(&data, sqlStr, attrId); err != nil {
		return data, err
	}
	return
}

func (a *attrDao) DeleteAttr(attrId int64) (err error) {
	tx, err := db.Beginx() // 开启事务
	if err != nil {
		fmt.Printf("begin trans failed, err:%v\n", err)
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
	// 删除 Attr
	sqlStr1 := `DELETE FROM attr WHERE attr_id = ?`
	rs, err := tx.Exec(sqlStr1, attrId)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	// 删除 AttrValue
	sqlStr2 := `DELETE FROM attr_value WHERE attr_id = ?`
	rs, err = tx.Exec(sqlStr2, attrId)
	if err != nil {
		return err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return err
	}

	return
}
