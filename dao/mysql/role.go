package mysql

import (
	"fmt"
	"math"
	"vue3_admin/model"
)

var RoleDao roleDao

type roleDao struct {
}

func (r *roleDao) CheckRoleExist(roleName string) (err error) {
	sqlStr := `SELECT count(role_id) FROM role WHERE role_name = ?`
	var count int
	if err = db.Get(&count, sqlStr, roleName); err != nil {
		return err
	}
	if count > 0 {
		return ErrorRoleExist
	}

	return
}

func (r *roleDao) InsertRole(role *model.Role) (err error) {
	sqlStr := `INSERT INTO role (role_id, role_name, remark) VALUES (?,?,?)`
	_, err = db.Exec(sqlStr, role.RoleID, role.RoleName, role.Remark)
	return err
}

func (r *roleDao) GetRoleList(roleName string, page, limit int64) (data *model.ResponseRoleList, err error) {
	countSqlStr := `SELECT COUNT(*) FROM role WHERE role_name like CONCAT('%',?,'%')`
	var count int64
	if err = db.Get(&count, countSqlStr, roleName); err != nil {
		return nil, err
	}
	sqlStr := `SELECT role_id, role_name, remark, create_time, update_time FROM role WHERE role_name like CONCAT('%',?,'%') LIMIT ?,?`
	roleList := make([]*model.Role, 0, 2)
	if err = db.Select(&roleList, sqlStr, roleName, (page-1)*limit, limit); err != nil {
		return nil, err
	}
	for _, role := range roleList {
		fmt.Printf("roleList: %#v\n", role)
	}

	data = &model.ResponseRoleList{
		Records:     roleList,
		Total:       count,
		Size:        limit,
		Current:     page,
		Pages:       int64(math.Ceil(float64(count) / float64(limit))),
		SearchCount: true,
	}
	return
}

func (*roleDao) UpdateRole(role *model.Role) (err error) {
	sqlStr := `UPDATE role SET role_name = ?, remark = ? WHERE role_id = ?`
	_, err = db.Exec(sqlStr, role.RoleName, role.Remark, role.RoleID)
	return err
}

func (*roleDao) DeleteRole(roleId int64) (err error) {
	sqlStr := `DELETE FROM role WHERE role_id = ?`
	_, err = db.Exec(sqlStr, roleId)
	return err
}

func (*roleDao) GetAllRoleList() (data []*model.Role, err error) {
	sqlStr := `SELECT role_id, role_name, remark, create_time, update_time FROM role`
	data = make([]*model.Role, 0, 2)
	if err = db.Select(&data, sqlStr); err != nil {
		return nil, err
	}
	return data, nil

}
