package main

import (
	e "github.com/pkg/errors"
)

type UserService struct {
}

func (s *UserService) GetNameById(id int) (string, error) {
	dao := new(UserDao)
	name, err := dao.GetNameById(id)
	if e.Is(e.Cause(err), ErrDbNothing) {
		// ... 一些处理
	}

	return name, err
}
