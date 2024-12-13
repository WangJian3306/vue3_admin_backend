package service

import (
	"go.uber.org/zap"
	"vue3_admin/dao/mysql"
	"vue3_admin/model"
)

var CategoryService categoryService

type categoryService struct {
}

func (c *categoryService) GetCategory1() (data []*model.Category1, err error) {
	data, err = mysql.CategoryDao.GetCategory1List()
	if err != nil {
		zap.L().Error("mysql.CategoryDao.GetCategory1List() failed", zap.Error(err))
		return
	}

	return
}

func (c *categoryService) GetCategory2(category1Id int64) (data []*model.Category2, err error) {
	data, err = mysql.CategoryDao.GetCategory2List(category1Id)
	if err != nil {
		zap.L().Error("mysql.CategoryDao.GetCategory2List() failed", zap.Error(err))
		return
	}

	return
}

func (c *categoryService) GetCategory3(category2Id int64) (data []*model.Category3, err error) {
	data, err = mysql.CategoryDao.GetCategory3List(category2Id)
	if err != nil {
		zap.L().Error("mysql.CategoryDao.GetCategory3List() failed", zap.Error(err))
		return
	}

	return
}

func (c *categoryService) CreateCategory2(p *model.ParamC2Create) (err error) {
	category2 := &model.Category2{
		Category2ID: p.Category2ID,
		Name:        p.Name,
		Category1ID: p.Category1ID,
	}

	err = mysql.CategoryDao.InsertCategory2(category2)

	return err
}

func (c *categoryService) CreateCategory3(p *model.ParamC3Create) (err error) {
	category3 := &model.Category3{
		Category3ID: p.Category3ID,
		Name:        p.Name,
		Category2ID: p.Category2ID,
	}

	err = mysql.CategoryDao.InsertCategory3(category3)

	return err
}
