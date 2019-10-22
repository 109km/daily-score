package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

var (
	UserList []*User
)

func GetUser(id int) (user User, err error) {

	res := User{Id: id}
	resErr := OrmInstance.Read(&res)

	if resErr == orm.ErrNoRows {
		return res, errors.New("查询不到")
	} else if resErr == orm.ErrMissPK {
		return res, errors.New("找不到主键")
	} else {
		return res, nil
	}

	return res, errors.New("User not exists")
}

func GetAllUsers() []*User {
	qs := OrmInstance.QueryTable("user")
	qs.All(&UserList)
	return UserList
}
