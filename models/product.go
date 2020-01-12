package models

import "time"

type Product struct {
	Id int
	Name string  `orm:"size(128)"`
	Link string  `orm:"size(256)"`
	Description string  `orm:"size(512)"`
	User  *User  `orm:"rel(fk)"`    //设置一对多关系（外键）
	Enable bool `orm:"default(true)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

