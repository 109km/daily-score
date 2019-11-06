package models

var (
	UserList      []*User
	responseError error
)

// Get single `user` by id
func GetUserById(id int64) (user User, errorMsg error) {
	user = User{Id: id}
	responseError = OrmInstance.Read(&user)
	errorMsg = ProceedSearchError(responseError)
	return user, errorMsg
}

// Get all users
func GetAllUsers() ([]*User, error) {
	qs := OrmInstance.QueryTable("user")
	_, responseError := qs.All(&UserList, "Id", "Mobile", "Nickname")

	if responseError != nil {
		return nil, ProceedSearchError(responseError)
	}

	return UserList, nil
}

func AddOneUser(mobile string, password string, nickname string) (uid int64, errorMsg error) {
	var user User
	var id int64
	user.Mobile = mobile
	user.Password = password
	user.Nickname = nickname

	id, responseError = OrmInstance.Insert(&user)
	errorMsg = ProceedSearchError(responseError)
	if responseError == nil {
		return id, nil
	}
	return -1, errorMsg

}

func LoginUser(mobile string, password string) (uid int64, errorMsg error) {
	var user User
	user.Mobile = mobile
	user.Password = password

	responseError = OrmInstance.Read(&user, "Mobile", "Password")
	errorMsg = ProceedSearchError(responseError)

	if responseError == nil {
		return user.Id, nil
	}
	return -1, errorMsg
}
