package model

import "time"

type BaseModel struct {
	ID         int64      `json:"ID,omitempty" db:"id"`
	CreateTime *time.Time `json:"createTime,omitempty" db:"create_time"`
	UpdateTime *time.Time `json:"updateTime,omitempty" db:"update_time"`
}
