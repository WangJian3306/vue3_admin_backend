package model

import "errors"

// 自定义 Error
var CodeMenuNodeExistError = errors.New("该节点下有子节点，不可以删除")
