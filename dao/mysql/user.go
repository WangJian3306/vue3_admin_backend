package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"math"
	"vue3_admin/model"
)

const secret = "devops"

var UserDao userDao

type userDao struct {
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	sum := h.Sum([]byte(oPassword))
	return hex.EncodeToString(sum)
}

func (u *userDao) Login(user *model.User) (err error) {
	oPassword := user.Password // 用户登录的密码
	sqlStr := `SELECT user_id, username, password FROM user WHERE username = ?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	// 查询数据库失败
	if err != nil {
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}

	return err
}

// GetUserById 根据id获取用户信息
func (u *userDao) GetUserById(uid int64) (user *model.User, err error) {
	user = new(model.User)
	sqlStr := `SELECT user_id,username, avatar from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	return
}

// CheckUserExist 检查指定用户名的用户是否存在
func (u *userDao) CheckUserExist(username string) (err error) {
	sqlStr := `SELECT count(user_id) FROM user WHERE username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}

	if count > 0 {
		return ErrorUserExist
	}

	return

}

// InsertUser 向数据库中插入一条新的用户记录
func (u *userDao) InsertUser(user *model.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)

	// 执行SQL语句入库
	sqlStr := `INSERT INTO user(user_id, username, name, password) VALUES(?,?,?,?) `
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Name, user.Password)

	return err
}

func (u *userDao) GetUserList(username string, page, limit int64) (data *model.ResponseUserList, err error) {
	countSqlStr := `SELECT count(user_id) FROM user`
	var count int64
	if err = db.Get(&count, countSqlStr); err != nil {
		return nil, err
	}
	sqlStr := `SELECT u.user_id, u.username, u.password, u.name, IFNULL(r.role_name, '') AS role_name, u.create_time, u.update_time FROM user u LEFT JOIN user_role ur on u.user_id = ur.user_id LEFT JOIN role r on ur.role_id = r.role_id WHERE u.username like CONCAT('%',?,'%') LIMIT ?,?;`
	userList := make([]*model.ResponseUser, 0, 2)
	if err = db.Select(&userList, sqlStr, username, (page-1)*limit, limit); err != nil {
		return nil, err
	}

	data = &model.ResponseUserList{
		Records: userList,
		Total:   count,
		Size:    limit,
		Current: page,
		Pages:   int64(math.Ceil(float64(count) / float64(limit))),
	}
	return
}

func (u *userDao) UpdateUser(user *model.User) (err error) {
	user.Password = encryptPassword(user.Password)
	sqlStr := `UPDATE user SET username=?, name=? WHERE user_id=?`
	_, err = db.Exec(sqlStr, user.Username, user.Name, user.UserID)
	return err
}

func (u *userDao) DeleteUser(userId int64) (err error) {
	sqlStr := `DELETE FROM user WHERE user_id=?`
	_, err = db.Exec(sqlStr, userId)
	return err
}

func (u *userDao) GetAssignRole(userId int64) (roleList []*model.Role, err error) {
	sqlStr := `SELECT r.role_id, r.role_name, r.remark,r.create_time, r.update_time FROM user_role as ur JOIN role r on ur.role_id = r.role_id WHERE ur.user_id = ?;`
	roleList = make([]*model.Role, 0, 2)
	if err = db.Select(&roleList, sqlStr, userId); err != nil {
		return nil, err
	}
	//fmt.Println(userId)
	//fmt.Printf("%#v\n", roleList)
	return roleList, nil
}

func (u *userDao) DeleteAssignRoleByUserId(userId int64) (err error) {
	sqlStr := `DELETE FROM user_role as ur WHERE ur.user_id = ?`
	_, err = db.Exec(sqlStr, userId)
	return err
}

func (u *userDao) DoAssign(p *model.ParamDoAssignRole) (err error) {
	sqlStr := `INSERT INTO user_role(user_id, role_id) VALUES(?,?)`
	for _, roleId := range p.RoleIDList {
		_, err = db.Exec(sqlStr, p.UserID, roleId)
		return err
	}
	return err
}

func (u *userDao) GetAssignMenuByUserId(userId int64) (menu []*model.Menu, err error) {
	sqlStr := `SELECT DISTINCT m.* FROM menu m INNER JOIN role_menu rm on m.menu_id = rm.menu_id INNER JOIN user_role ur on rm.role_id = ur.role_id WHERE ur.user_id = ?;`
	if err = db.Select(&menu, sqlStr, userId); err != nil {
		return nil, err
	}
	return menu, nil
}
