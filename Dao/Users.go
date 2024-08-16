package Dao

import (
	"errors"
	"goVben/Models"
)

func GetLoginUser(user *Models.Users) error {
	err := DB.Where("account = ? and password = ?", user.Account, user.Password).First(&user).Error
	if err != nil {
		return errors.New("请检查用户名与密码")
	}
	return nil
}
