package service

import (
	"go.uber.org/zap"
	"vue3_admin/dao/mysql"
	"vue3_admin/model"
	"vue3_admin/pkg/snowflake"
)

var TrademarkService trademarkService

type trademarkService struct {
}

func (t *trademarkService) CreateTrademark(p *model.ParamTmSave) (err error) {
	// 1. 判断品牌是否已经存在
	if err = mysql.TrademarkDao.CheckTrademarkExist(p.TmName); err != nil {
		return err
	}
	// 2. 生成 TmID
	TmID := snowflake.GenID()

	// 3. 构造一个品牌实例
	trademark := &model.Trademark{
		TmID:    TmID,
		TmName:  p.TmName,
		LogoUrl: p.LogoUrl,
	}

	// 4. 保存进数据库
	err = mysql.TrademarkDao.InsertTrademark(trademark)

	return err
}

func (t *trademarkService) GetTrademarkList(page, limit int64) (data *model.ResponseTmList, err error) {
	data, err = mysql.TrademarkDao.GetTrademarkList(page, limit)
	if err != nil {
		zap.L().Error("mysql.TrademarkDao.GetTrademarkList() failed", zap.Error(err))
		return
	}

	return
}

func (t *trademarkService) UpdateTrademark(p *model.ParamTmUpdate) (err error) {
	// 构造一个品牌实例
	trademark := &model.Trademark{
		TmID:    p.TmID,
		TmName:  p.TmName,
		LogoUrl: p.LogoUrl,
	}

	// 保存进数据库
	err = mysql.TrademarkDao.UpdateTrademark(trademark)

	return err

}

func (t *trademarkService) DeleteTrademark(tmId int64) (err error) {
	return mysql.TrademarkDao.DeleteTrademark(tmId)
}
