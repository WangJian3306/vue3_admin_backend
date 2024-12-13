package model

type Category1 struct {
	BaseModel

	CategoryID int64  `json:"id" db:"category1_id"`
	Name       string `json:"name" db:"name"`
}

type Category2 struct {
	BaseModel

	Category2ID int64  `json:"id" db:"category2_id"`
	Name        string `json:"name" db:"name"`
	Category1ID int64  `json:"category1Id" db:"category1_id"`
}

type Category3 struct {
	BaseModel

	Category3ID int64  `json:"id" db:"category3_id"`
	Name        string `json:"name" db:"name"`
	Category2ID int64  `json:"category2Id" db:"category2_id"`
}

type ParamC2Create struct {
	Category2ID int64  `json:"category2Id"`
	Name        string `json:"name"`
	Category1ID int64  `json:"category1Id"`
}

type ParamC3Create struct {
	Category3ID int64  `json:"category3Id"`
	Name        string `json:"name"`
	Category2ID int64  `json:"category2Id"`
}
