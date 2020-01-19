package models

import (
	"fmt"
	"productgo/utils"
	"time"
)

type Product struct {
	Id int
	Name string  `orm:"size(128)"`
	Link string  `orm:"size(256)"`
	Description string  `orm:"size(512)"`
	User  *User  `orm:"rel(fk)"`    //设置一对多关系（外键）
	VoteCount int `orm:"default(0)"`
	Enable bool `orm:"default(true)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func GetProductById(id int) (*Product, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid id: %d", id)
	}
	o := utils.GetOrmer()
	product := Product{Id: id}
	if err := o.Read(&product); err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) Vote(user *User) error {
	// 首先判断是否已经点过赞
	if !p.HasVoted(user.Id) {
		// 创建点赞记录
		pv := ProductVote{
			User: user,
			Product: p,
		}
		o := utils.GetOrmer()
		_, err := o.Insert(&pv)
		if err != nil {
			return err
		}
		// 给product点赞数加1
		p.VoteCount += 1
		_, err = o.Update(p)
		//TODO：事务处理
		return err
	}
	return nil
}

func (p *Product) HasVoted(uid int) bool {
	if uid <= 0 {
		return false
	}
	o := utils.GetOrmer()
	// 首先判断是否已经点过赞
	qs := o.QueryTable("product_vote")
	return qs.Filter("user_id", uid).Filter("product_id", p.Id).Exist()
}

//product_vote
type ProductVote struct {
	Id int
	User *User `orm:"rel(fk)"`
	Product *Product `orm:"rel(fk)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}
