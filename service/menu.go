package service

import (
	"vue3_admin/dao/mysql"
	"vue3_admin/model"
	"vue3_admin/pkg/menuhelper"
	"vue3_admin/pkg/snowflake"
)

var MenuService menuService

type menuService struct{}

func (*menuService) GetMenu() (data []model.Menu, err error) {
	// 1. 查询所有菜单
	menuList, err := mysql.MenuDao.GetMenuList()
	if err != nil {
		return nil, err
	}
	if len(menuList) == 0 {
		return nil, nil
	}

	// 2. 格式化成返回需要的树形格式
	data, err = menuhelper.BuildTree(menuList)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (*menuService) SaveMenu(p *model.ParamMenuSave) (err error) {
	// 1. 判断菜单是否已经存在
	if err = mysql.MenuDao.CheckMenuExist(p.Name); err != nil {
		return err
	}

	menuId := snowflake.GenID()
	// 2. 生成一个菜单实例
	menu := &model.Menu{
		MenuID: menuId,
		Name:   p.Name,
		PID:    p.PID,
		CODE:   p.CODE,
		TOCODE: "",
		TYPE:   p.TYPE,
		STATUS: "",
		LEVEL:  p.LEVEL,
		SELECT: false,
	}
	// 3. 保存进数据库
	err = mysql.MenuDao.InsertMenu(menu)

	return err
}

func (*menuService) UpdateMenu(p *model.ParamMenuUpdate) (err error) {
	// 构造一个菜单实例
	menu := &model.Menu{
		MenuID: p.MenuID,
		Name:   p.Name,
		PID:    p.PID,
		CODE:   p.CODE,
		LEVEL:  p.LEVEL,
	}

	// 保存进数据库
	err = mysql.MenuDao.UpdateMenu(menu)

	return err
}

func (*menuService) DeleteMenu(menuId int64) (err error) {
	// 根据当前菜单ID，查询是否包含子菜单
	count, err := mysql.MenuDao.QuerySubMenuByID(menuId)
	if err != nil {
		return err
	}
	if count > 0 {
		return model.CodeMenuNodeExistError
	}

	// 根据菜单ID删除菜单
	err = mysql.MenuDao.DeleteMenuByID(menuId)

	return err
}

func (*menuService) ToAssign(roleId int64) (data []model.Menu, err error) {
	// 1. 查询所有菜单
	menuList, err := mysql.MenuDao.GetMenuList()
	if err != nil {
		return nil, err
	}
	if len(menuList) == 0 {
		return nil, nil
	}

	// 2. 查询已经具有的权限
	assignMenuIdList, err := mysql.MenuDao.GetAssignMenu(roleId)
	if err != nil {
		return nil, err
	}

	// 3. 将所有权限中已经具有的权限 select 值置为 true
	if len(assignMenuIdList) != 0 {
		for i := 0; i < len(menuList); i++ {
			for _, menuId := range assignMenuIdList {
				if menuList[i].MenuID == menuId {
					menuList[i].SELECT = true
				}
			}
		}
	}

	// 3. 格式化成返回需要的树形格式
	data, err = menuhelper.BuildTree(menuList)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (*menuService) DoAssign(roleId int64, menuIds []int64) (err error) {
	// 根据角色ID删除已分配的菜单
	err = mysql.MenuDao.DeleteAssignMenuByRoleId(roleId)
	if err != nil {
		return err
	}

	// 再重新给角色分配菜单权限
	err = mysql.MenuDao.DoAssign(roleId, menuIds)

	return err
}
