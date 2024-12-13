package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"vue3_admin/model"
)

var CategoryDao categoryDao

type categoryDao struct{}

func (c *categoryDao) GetCategory1List() (data []*model.Category1, err error) {
	sqlStr := `SELECT category1_id, name FROM category1`
	if err = db.Select(&data, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no category1 in db")
			return
		}
	}

	return
}

func (c *categoryDao) GetCategory2List(category1Id int64) (data []*model.Category2, err error) {
	sqlStr := `SELECT category2_id, name, category1_id FROM category2 WHERE category1_id = ?`
	if err = db.Select(&data, sqlStr, category1Id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no category2 in db")
			return
		}
	}

	return
}

func (c *categoryDao) GetCategory3List(category2Id int64) (data []*model.Category3, err error) {
	sqlStr := `SELECT category3_id, name, category2_id FROM category3 WHERE category2_id = ?`
	if err = db.Select(&data, sqlStr, category2Id); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no category3 in db")
			return
		}
	}

	return
}

func (c *categoryDao) InsertCategory2(data *model.Category2) (err error) {
	sqlStr := `INSERT INTO category2 (category2_id, name, category1_id) VALUES (?,?,?)`
	_, err = db.Exec(sqlStr, data.Category2ID, data.Name, data.Category1ID)

	return err
}

func (c *categoryDao) InsertCategory3(data *model.Category3) (err error) {
	sqlStr := `INSERT INTO category3 (category3_id, name, category2_id) VALUES (?,?,?)`
	_, err = db.Exec(sqlStr, data.Category3ID, data.Name, data.Category2ID)

	return err
}
