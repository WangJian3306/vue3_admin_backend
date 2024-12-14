package mysql

import (
	"math"
	"vue3_admin/model"
)

var TrademarkDao trademarkDao

type trademarkDao struct {
}

func (*trademarkDao) CheckTrademarkExist(tmName string) (err error) {
	sqlStr := `SELECT count(tm_id) from trademark WHERE tm_name = ?`
	var count int
	if err = db.Get(&count, sqlStr, tmName); err != nil {
		return err
	}
	if count > 0 {
		return ErrorTrademarkExist
	}

	return
}

func (*trademarkDao) InsertTrademark(tm *model.Trademark) (err error) {

	// 执行SQL 语句入库
	sqlStr := `INSERT INTO trademark (tm_id, tm_name, logo_url) VALUES (?,?,?)`
	_, err = db.Exec(sqlStr, tm.TmID, tm.TmName, tm.LogoUrl)

	return err
}

func (*trademarkDao) GetTrademarkList(page, limit int64) (data *model.ResponseTmList, err error) {
	countSqlStr := `SELECT count(tm_id) from trademark`
	var count int64
	if err = db.Get(&count, countSqlStr); err != nil {
		return nil, err
	}
	sqlStr := `SELECT tm_id, tm_name, logo_url, create_time, update_time FROM trademark LIMIT ?,?`
	trademarks := make([]*model.Trademark, 0, 2)
	if err = db.Select(&trademarks, sqlStr, (page-1)*limit, limit); err != nil {
		return nil, err
	}
	data = &model.ResponseTmList{
		Records:     trademarks,
		Total:       count,
		Size:        limit,
		Current:     page,
		SearchCount: true,
		Pages:       int64(math.Ceil(float64(count) / float64(limit))),
	}

	return
}

func (*trademarkDao) UpdateTrademark(data *model.Trademark) (err error) {
	sqlStr := `UPDATE trademark SET tm_name = ?, logo_url = ? WHERE tm_id = ?`
	_, err = db.Exec(sqlStr, data.TmName, data.LogoUrl, data.TmID)

	return err
}

func (*trademarkDao) DeleteTrademark(tmId int64) (err error) {
	sqlStr := `DELETE FROM trademark WHERE tm_id = ?`
	_, err = db.Exec(sqlStr, tmId)
	return err
}

func (*trademarkDao) GetAllTrademarkList() (data []model.Trademark, err error) {
	sqlStr := `SELECT tm_id, tm_name, logo_url FROM trademark`
	if err = db.Select(&data, sqlStr); err != nil {
		return nil, err
	}
	return data, nil
}
