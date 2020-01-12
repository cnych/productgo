package models

import (
	"fmt"
	"productgo/utils"
	"time"
)

type User struct {
	Id         int
	Username   string `orm:"size(64)"`
	Password   string `orm:"size(64)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func InsertUser(user *User) (int64, error) {
	sql := "insert into user(username, password) values(?, ?)"
	return utils.ExecSQL(sql, user.Username, user.Password)
}

//1 a 12345
//2 b 12345
//a 12345
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("select id from user where username='%s'", username)
	row := utils.QueryRow(sql)

	id := 0
	_ = row.Scan(&id)

	return id
}

func QueryUserWithParam(username, password string) *User {
	sql := fmt.Sprintf("select id from user where username='%s' and password='%s'", username, password)
	row := utils.QueryRow(sql)

	id := 0
	_ = row.Scan(&id)

	return &User{Id: id, Username: username, Password: password}
}
