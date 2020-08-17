package accountEnums

import (
	enumsbase "grpc-demo/enums"
)

const (
	RoleUser       = 1    // 普通用户
	RoleAdmin      = 1024 // 系统管理员
)

func NewRoleEnums() enumsbase.EnumBaseInterface {
	return enumsbase.EnumBase{
		Enums: []int{1, 1024},
	}
}
