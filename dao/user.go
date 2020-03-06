package dao

import (
	"productgo/commons"
	"productgo/commons/utils"
	"productgo/models"
	"time"
)

// GetUserByParam 根据用户名和密码查询用户
func GetUserByParam(username, password string) (*models.User, error) {
	var user models.User
	qs := commons.GetOrmer().QueryTable("user")
	err := qs.Filter("username", username).Filter("password", utils.MD5(password)).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UserExists(username string) bool {
	qs := commons.GetOrmer().QueryTable("user")
	return qs.Filter("username", username).Exist()
}

// AddUser 新建用户
func AddUser(user *models.User) (int64, error) {
	now := time.Now()
	user.Password = utils.MD5(user.Password)
	user.Created = now
	id, err := commons.GetOrmer().Insert(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
