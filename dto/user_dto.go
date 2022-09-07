package dto

import "xukun/model"

/*
处理用户返回时的敏感字段
 */

type UserDto struct {
	Name string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto  {
	return UserDto{
		Name: user.Name,
		Telephone: user.Telephone,
	}
}