package service

import (
	"go.uber.org/zap"
	"vue3_admin/dao/mysql"
	"vue3_admin/model"
	"vue3_admin/pkg/snowflake"
)

var AttrService attrService

type attrService struct {
}

func (a *attrService) UpdateAttr(p *model.Attr) (err error) {
	err = mysql.AttrDao.UpdateAttrAndAttrValue(p)
	if err != nil {
		zap.L().Error("mysql.AttrDao.UpdateAttrAndAttrValue() failed", zap.Error(err))
		return
	}
	return err
}

func (a *attrService) CreateAttr(p *model.Attr) (err error) {
	// 创建属性实例
	attr := &model.Attr{
		AttrName:      p.AttrName,
		CategoryId:    p.CategoryId,
		AttrID:        snowflake.GenID(),
		CategoryLevel: p.CategoryLevel,
	}

	// 创建属性值实例
	//var attrValues []*model.AttrValue
	for i := 0; i < len(p.AttrValueList); i++ {
		newAttrValue := &model.AttrValue{
			ValueName:   p.AttrValueList[i].ValueName,
			AttrID:      attr.AttrID,
			AttrValueID: snowflake.GenID(),
		}
		attr.AttrValueList = append(attr.AttrValueList, newAttrValue)
	}

	// 写入数据库
	err = mysql.AttrDao.InsertAttrAndAttrValue(attr)
	if err != nil {
		zap.L().Error("mysqlAttrDao.InsertAttrAndAttrValue() failed", zap.Error(err))
		return
	}

	//err = mysql.AttrDao.InsertAttrValue(attrValues)
	//if err != nil {
	//	zap.L().Error("AttrDao.InsertAttrValue() failed", zap.Error(err))
	//	return
	//}

	return err
}

func (a *attrService) GetAttr(c1Id, c2Id, c3Id int64) (attrs []*model.Attr, err error) {

	attrs, err = mysql.AttrDao.GetAttr(c1Id, c2Id, c3Id)
	if err != nil {
		zap.L().Error("mysql.AttrDao.GetAttr() failed", zap.Error(err))
		return
	}

	var attrValueList []*model.AttrValue
	for _, attr := range attrs {
		attrValueList, err = mysql.AttrDao.GetAttrValue(attr.AttrID)
		if err != nil {
			zap.L().Error("mysql.AttrDao.GetAttrValue() failed", zap.Error(err))
			return
		}
		attr.AttrValueList = attrValueList
	}
	return
}

func (a *attrService) DeleteAttr(attrId int64) (err error) {
	return mysql.AttrDao.DeleteAttr(attrId)
}
