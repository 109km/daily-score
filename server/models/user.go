package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

var (
	UserList []*User
	opManager orm.Ormer
)


func init() {
	opManager = orm.NewOrm()
}


func GetUser(uid int) (user User, err error) {
	
	resUser := User{Id: uid}
	resErr := opManager.Read(&resUser)

	if resErr == orm.ErrNoRows {
		return resUser, errors.New("查询不到")
	}else if resErr == orm.ErrMissPK {
		return resUser, errors.New("找不到主键")
	}else {
		return resUser, nil
	}
	
	return resUser, errors.New("User not exists")
}

func GetAllUsers() []*User {
	qs := opManager.QueryTable("user")
	qs.All(&UserList)
	return UserList
}
