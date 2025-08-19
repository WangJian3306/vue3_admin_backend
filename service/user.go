package service

import (
	"fmt"
	"go.uber.org/zap"
	"vue3_admin/dao/mysql"
	"vue3_admin/model"
	"vue3_admin/pkg/jwt"
	"vue3_admin/pkg/snowflake"
)

var UserService userService

type userService struct {
}

func (*userService) Login(p *model.ParamUserLogin) (token string, err error) {
	user := &model.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能找到user.UserID
	if err = mysql.UserDao.Login(user); err != nil {
		return "", err
	}
	// 生成 JWT
	token, err = jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return "", err
	}
	return token, err
}

func (*userService) SignUp(p *model.ParamUserSignUp) (err error) {
	// 1. 判断用户是否存在
	if err = mysql.UserDao.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2. 生成UID
	userID := snowflake.GenID()

	// 构造一个User实例
	user := &model.User{
		UserID:   userID,
		Username: p.Username,
		Name:     p.Name,
		Password: p.Password,
	}

	// 3. 保存进数据库
	err = mysql.UserDao.InsertUser(user)

	return err
}

func (*userService) GetUserInfo(userID int64) (*model.ResponseUserInfo, error) {
	user, err := mysql.UserDao.GetUserById(userID)
	if err != nil {
		zap.L().Error("mysql.UserDao.GetUserById(userID) failed",
			zap.Int64("user_id", userID),
			zap.Error(err))
	}
	// 获取用户角色
	roleList, err := mysql.UserDao.GetAssignRole(userID)
	if err != nil {
		zap.L().Error("mysql.UserDao.GetAssignRole(userID) failed", zap.Error(err))
		return nil, err
	}
	roleNameList := make([]string, 0)
	for _, role := range roleList {
		roleNameList = append(roleNameList, role.RoleName)
	}

	// 获取当前用户拥有的所有菜单
	menuList, err := mysql.UserDao.GetAssignMenuByUserId(userID)
	if err != nil {
		zap.L().Error("mysql.UserDao.GetAssignMenuByUserId(userID) failed", zap.Error(err))
		return nil, err
	}
	// 获取用户路由 和 用户按钮权限
	userRoutes := make([]string, 0, 12)
	userButtons := make([]string, 0, 12)
	for _, menu := range menuList {
		if menu.LEVEL != 4 {
			userRoutes = append(userRoutes, menu.CODE)
		} else {
			userButtons = append(userButtons, menu.CODE)
		}
	}
	fmt.Printf("userRoutes: %#v\n", userRoutes)
	fmt.Printf("userButtons: %#v\n", userButtons)

	userInfo := &model.ResponseUserInfo{
		Routes:  userRoutes,
		Buttons: userButtons,
		Roles:   roleNameList,
		Name:    user.Username,
		Avatar:  user.Avatar,
	}
	return userInfo, err

}

func (*userService) GetUserList(username string, page, limit int64) (data *model.ResponseUserList, err error) {
	data, err = mysql.UserDao.GetUserList(username, page, limit)
	if err != nil {
		zap.L().Error("mysql.UserDao.GetUserList() failed", zap.Error(err))
		return
	}

	return
}

func (*userService) UpdateUser(p *model.ParamUserUpdate) (err error) {
	user := &model.User{
		UserID:   p.UserID,
		Username: p.Username,
		Name:     p.Name,
	}
	err = mysql.UserDao.UpdateUser(user)
	return err

}

func (*userService) DeleteUser(userId int64) (err error) {
	return mysql.UserDao.DeleteUser(userId)
}

func (*userService) ToAssign(userId int64) (data *model.ResponseToAssignRole, err error) {
	roleList, err := mysql.RoleDao.GetAllRoleList()
	if err != nil {
		zap.L().Error("mysql.RoleDao.GetAllRoleList() failed", zap.Error(err))
		return nil, err
	}
	assignRoleList, err := mysql.UserDao.GetAssignRole(userId)
	if err != nil {
		zap.L().Error("mysql.UserDao.GetAssignRole() failed", zap.Error(err))
		return nil, err
	}
	toAssign := &model.ResponseToAssignRole{
		AssignRoles:  assignRoleList,
		AllRolesList: roleList,
	}
	return toAssign, err
}

func (*userService) DoAssignRole(p *model.ParamDoAssignRole) (err error) {
	// 1. 根据用户 ID 删除所有的角色分配
	err = mysql.UserDao.DeleteAssignRoleByUserId(p.UserID)
	if err != nil {
		zap.L().Error("mysql.UserDao.DeleteAssignRoleByUserId() failed", zap.Error(err))
		return err
	}

	// 2. 再重新给用户分配角色
	err = mysql.UserDao.DoAssign(p)
	return err
}
