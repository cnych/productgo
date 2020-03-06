package models

import (
	"strings"
	"time"
)

type User struct {
	Id       int
	Username string    `orm:"size(64)"`
	Password string    `orm:"size(64)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
}

type Avatar struct {
	L string // 颜色
	C string // 内容
}

func (u *User) MkAvatar() *Avatar {
	colors := []string{
		"#1abc9c",
		"#2ecc71",
		"#3498db",
		"#1abc9c",
		"#2ecc71",
		"#3498db",
		"#1abc9c",
		"#2ecc71",
		"#3498db",
		"#2c3e50",
	}
	index := u.Id % 10
	return &Avatar{
		L: colors[index],
		C: strings.ToUpper(u.Username[:1]),
	}
}
