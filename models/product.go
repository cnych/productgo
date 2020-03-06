package models

import (
	"productgo/commons"
	"time"
)

type Product struct {
	Id          int64
	Name        string    `orm:"size(128)"`
	Link        string    `orm:"size(256)"`
	Description string    `orm:"size(512)"`
	User        *User     `orm:"rel(fk)"` //设置一对多关系（外键）
	VoteCount   int       `orm:"default(0)"`
	Enable      bool      `orm:"default(true)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
}

func (p *Product) HasVoted(uid int) bool {
	if uid <= 0 {
		return false
	}
	// 首先判断是否已经点过赞
	qs := commons.GetOrmer().QueryTable("product_vote")
	return qs.Filter("user_id", uid).Filter("product_id", p.Id).Exist()
}

//product_vote
type ProductVote struct {
	Id      int
	User    *User     `orm:"rel(fk)"`
	Product *Product  `orm:"rel(fk)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

type ProductDateItem struct {
	Date     string
	Products []Product
	Uid      int
}
