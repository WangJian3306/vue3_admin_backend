package model

type Attr struct {
	BaseModel

	AttrID        int64        `json:"id" db:"attr_id"`                                      // 属性ID
	AttrName      string       `json:"attrName" db:"attr_name" binding:"required"`           // 属性名称
	CategoryId    int64        `json:"categoryId" db:"category_id" binding:"required"`       // 三级分类ID
	CategoryLevel int          `json:"categoryLevel" db:"category_level" binding:"required"` // 分类级别
	AttrValueList []*AttrValue `json:"attrValueList"`                                        // 属性值列表
}

type AttrValue struct {
	AttrValueID int64  `json:"id" db:"attr_value_id"`
	ValueName   string `json:"valueName" db:"value_name" binding:"required"`
	AttrID      int64  `json:"attrId" db:"attr_id"`
}

type ParamAttrCreate struct {
	AttrID        int64                   `json:"id"`
	AttrName      string                  `json:"attrName" binding:"required"`
	CategoryId    int64                   `json:"categoryId" binding:"required"`
	CategoryLevel int                     `json:"categoryLevel" binding:"required"`
	AttrValueList []*ParamAttrValueCreate `json:"attrValueList"` // 属性值列表
}

type ParamAttrValueCreate struct {
	ValueName string `json:"valueName" binding:"required"`
}

//func GetCommunityList() (communityList []*models.Community, err error) {
//	sqlStr := `SELECT community_id, community_name FROM community`
//	if err = db.Select(&communityList, sqlStr); err != nil {
//		if err == sql.ErrNoRows {
//			zap.L().Warn("there is no community in db")
//			return nil, err
//		}
//	}
//
//	return
//}
