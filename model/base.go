package model

import (
	"vue3_admin/pkg/SimpleDateFormat"
)

type BaseModel struct {
	ID         int64                              `json:"ID,omitempty" db:"id"`
	CreateTime *SimpleDateFormat.SimpleDateFormat `json:"createTime,omitempty" db:"create_time"`
	UpdateTime *SimpleDateFormat.SimpleDateFormat `json:"updateTime,omitempty" db:"update_time"`
}
