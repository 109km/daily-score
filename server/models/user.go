package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

var (
	UserList map[int]*User
)

func init() {
	UserList = make(map[int]*User)
}


func GetUser(uid int) (user User, err error) {
	o := orm.NewOrm()
	resUser := User{Id: uid}
	resErr := o.Read(&resUser)

	if resErr == orm.ErrNoRows {
		return resUser, errors.New("查询不到")
	}else if resErr == orm.ErrMissPK {
		return resUser, errors.New("找不到主键")
	}else {
		return resUser, nil
	}
	
	return resUser, errors.New("User not exists")
}

func GetAllUsers() map[int]*User {	
	return UserList
}
