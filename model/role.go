package model

type Role struct {
	BaseModel

	RoleID   int64  `db:"role_id" json:"id"`
	RoleName string `db:"role_name" json:"roleName"`
	Remark   string `db:"remark" json:"remark"`
}

type ParamRoleSave struct {
	RoleName string `json:"roleName"`         // 角色名称，不能为空
	Remark   string `json:"remark,omitempty"` // 角色备注
}

type ParamRoleUpdate struct {
	RoleID   int64  `json:"id"`               // 角色ID，
	RoleName string `json:"roleName"`         // 角色名称，不能为空
	Remark   string `json:"remark,omitempty"` // 角色备注
}

// ResponseRoleList 角色列表返回数据
type ResponseRoleList struct {
	Records     []*Role `json:"records"`
	Total       int64   `json:"total"`
	Size        int64   `json:"size"`
	Current     int64   `json:"current"`
	Pages       int64   `json:"pages"`
	SearchCount bool    `json:"searchCount"`
}

type ResponseToAssignRole struct {
	AssignRoles  []*Role `json:"assignRoles"`
	AllRolesList []*Role `json:"allRolesList"`
}
