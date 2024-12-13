package model

type Menu struct {
	BaseModel
	MenuID   int64  `db:"menu_id" json:"id"`
	Name     string `db:"name" json:"name"`
	PID      int64  `db:"pid" json:"pid"`
	CODE     string `db:"code" json:"code"`
	TOCODE   string `db:"to_code" json:"toCode"`
	TYPE     int    `db:"type" json:"type"`
	STATUS   string `db:"status" json:"status"`
	LEVEL    int    `db:"level" json:"level"`
	CHILDREN []Menu `db:"children" json:"children"`
	SELECT   bool   `db:"select" json:"select"`
}

type ParamMenuSave struct {
	Name  string `db:"name" json:"name"`
	PID   int64  `db:"pid" json:"pid"`
	CODE  string `db:"code" json:"code"`
	TYPE  int    `db:"type" json:"type"`
	LEVEL int    `db:"level" json:"level"`
}

type ParamMenuUpdate struct {
	MenuID int64  `db:"menu_id" json:"id"`
	Name   string `db:"name" json:"name"`
	PID    int64  `db:"pid" json:"pid"`
	CODE   string `db:"code" json:"code"`
	LEVEL  int    `db:"level" json:"level"`
}
