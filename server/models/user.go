package models

var (
	UserList []*User
)

// Get single `user` by id
func GetUserById(id int) (user User, err error) {
	res := User{Id: id}
	resErr := OrmInstance.Read(&res)

	errorMsg := ProceedSearchError(resErr)
	return res, errorMsg
}

// Get all users
func GetAllUsers() []*User {
	qs := OrmInstance.QueryTable("user")
	qs.All(&UserList)
	return UserList
}

func AddOneUser(mobile string, password string, nickname string) (uid int64, err error) {
	var user User
	user.Mobile = mobile
	user.Password = password
	user.Nickname = nickname

	id, err := OrmInstance.Insert(&user)
	if err == nil {
		return id, nil
	}
	return 0, err

}
