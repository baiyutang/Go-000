package main

import (
	"database/sql"
	"errors"
	e "github.com/pkg/errors"
	_ "github.com/siddontang/go-mysql/driver"
)

const (
	_dsn   = "root:root@127.0.0.1:3306?go"
	_table = "users"
)

var ErrDbNothing = errors.New("db:no rows in result set")

type UserDao struct {
	ErrNoRows error
}

func (d *UserDao) GetNameById(id int) (string, error) {
	db, _ := sql.Open("mysql", _dsn)
	defer db.Close()
	var name string
	err := db.QueryRow("select `name` from "+_table+" where id = ?", id).Scan(&name)
	// 这里判断找不到记录错误进行一次转换
	if e.Is(err, sql.ErrNoRows) {
		return name, ErrDbNothing
	}
	return name, nil
}
