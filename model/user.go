package model

type User struct {
	BaseModel

	UserID   int64  `db:"user_id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Name     string `db:"name" json:"name"`
	Phone    string `db:"phone" json:"phone"`
	Avatar   string `db:"avatar" json:"avatar,omitempty"`
}

// ParamUserLogin 用户登录参数
type ParamUserLogin struct {
	Username string `json:"username" binding:"required"` // 用户名，不能为空
	Password string `json:"password" binding:"required"` // 密码，不能为空
}

// ParamUserSignUp 注册用户参数
type ParamUserSignUp struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	//RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	//Avatar     string `json:"avatar" binding:"required"`
}

// ParamUserUpdate 更新用户参数
type ParamUserUpdate struct {
	UserID   int64  `json:"id" binding:"required"`       // 用户ID
	Username string `json:"username" binding:"required"` // 用户姓名
	Name     string `json:"name" binding:"required"`     // 用户昵称
}

// ResponseUserInfo 用户信息返回数据
type ResponseUserInfo struct {
	Routes  []string `json:"routes"`
	Buttons []string `json:"buttons"`
	Roles   []string `json:"roles"`
	Name    string   `json:"name"`
	Avatar  string   `json:"avatar"`
}

// ResponseUser 用户返回数据
type ResponseUser struct {
	BaseModel
	UserID   int64  `db:"user_id" json:"id"`
	Username string `db:"username" json:"username"`
	RoleName string `db:"role_name" json:"roleName"`
	Password string `db:"password" json:"password"`
	Name     string `db:"name" json:"name"`
	Phone    string `db:"phone" json:"phone"`
}

// ResponseUserList 用户列表返回数据
type ResponseUserList struct {
	Records []*ResponseUser `json:"records"`
	Total   int64           `json:"total"`
	Size    int64           `json:"size"`
	Current int64           `json:"current"`
	Pages   int64           `json:"pages"`
}

type ParamDoAssignRole struct {
	UserID     int64   `json:"userId"`
	RoleIDList []int64 `json:"roleIdList"`
}
