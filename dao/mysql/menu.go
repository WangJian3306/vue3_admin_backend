package mysql

import (
	"go.uber.org/zap"
	"vue3_admin/model"
)

var MenuDao menuDao

type menuDao struct{}

func (*menuDao) GetMenuList() (data []model.Menu, err error) {
	sqlStr := `SELECT menu_id, name, pid, code, to_code,type, status, level FROM menu`
	if err = db.Select(&data, sqlStr); err != nil {
		zap.L().Error("dao.menuDao.GetMenuList() error", zap.Error(err))
		return nil, err
	}
	nilMenu := make([]model.Menu, 0, 1)
	for _, menu := range data {
		if menu.CHILDREN == nil {
			menu.CHILDREN = nilMenu
		}
	}
	return data, err
}

func (*menuDao) CheckMenuExist(menuName string) (err error) {
	sqlStr := `SELECT COUNT(1) FROM menu WHERE name=?`
	var cnt int
	if err = db.Get(&cnt, sqlStr, menuName); err != nil {
		return err
	}
	if cnt > 0 {
		return ErrorMenuExist
	}
	return
}

func (*menuDao) InsertMenu(menu *model.Menu) (err error) {
	sqlStr := `INSERT INTO menu(menu_id,name, pid, code, to_code, type, status, level) VALUES (?, ?, ?, ?, ?, ?, ?, ? )`
	_, err = db.Exec(sqlStr, menu.MenuID, menu.Name, menu.PID, menu.CODE, menu.TOCODE, menu.TYPE, menu.STATUS, menu.LEVEL)
	return err
}

func (*menuDao) UpdateMenu(menu *model.Menu) (err error) {
	sqlStr := `UPDATE menu SET name = ?, pid = ?, code = ?, level = ? WHERE menu_id = ?`
	_, err = db.Exec(sqlStr, menu.Name, menu.PID, menu.CODE, menu.LEVEL, menu.MenuID)
	return err
}

func (*menuDao) QuerySubMenuByID(menuId int64) (count int64, err error) {
	sqlStr := `SELECT COUNT(1) FROM menu WHERE pid = ?`
	if err = db.Get(&count, sqlStr, menuId); err != nil {
		return count, err
	}
	return count, nil
}

func (*menuDao) DeleteMenuByID(menuId int64) (err error) {
	sqlStr := `DELETE FROM menu WHERE menu_id = ?`
	_, err = db.Exec(sqlStr, menuId)
	return err
}

func (*menuDao) GetAssignMenu(roleId int64) (menuIdList []int64, err error) {
	sqlStr := `SELECT menu_id FROM role_menu WHERE role_id = ?`
	if err = db.Select(&menuIdList, sqlStr, roleId); err != nil {
		return nil, err
	}
	return menuIdList, nil
}

func (*menuDao) DeleteAssignMenuByRoleId(roleId int64) (err error) {
	sqlStr := `DELETE FROM role_menu WHERE role_id = ?`
	_, err = db.Exec(sqlStr, roleId)
	return err
}

func (*menuDao) DoAssign(roleId int64, menuIds []int64) (err error) {
	sqlStr := `INSERT INTO role_menu (role_id, menu_id) VALUES (?,?)`
	for _, menuId := range menuIds {
		if _, err := db.Exec(sqlStr, roleId, menuId); err != nil {
			return err
		}
	}
	return err
}
