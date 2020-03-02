package dao

import (
	"fmt"
	"productgo/models"
	"productgo/utils"
	"time"
)

// AddProduct 添加产品
func AddProduct(product *models.Product) (int64, error) {
	now := time.Now()
	product.Created = now
	id, err := GetOrmer().Insert(product)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetProductById 根据ID获取产品信息
func GetProductById(id int64) (*models.Product, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid id: %d", id)
	}
	product := models.Product{Id: id}
	if err := GetOrmer().Read(&product); err != nil {
		return nil, err
	}
	return &product, nil
}

// Delete、List、Vote

func VoteProduct(user *models.User, product *models.Product) error {
	// 首先判断是否已经点过赞
	if !product.HasVoted(user.Id) {
		// 创建点赞记录
		pv := models.ProductVote{
			User:    user,
			Product: product,
		}
		_, err := GetOrmer().Insert(&pv)
		if err != nil {
			return err
		}
		// 给product点赞数加1
		product.VoteCount += 1
		_, err = GetOrmer().Update(product)
		//TODO：事务处理
		return err
	}
	return nil
}

func GetProductsByDate(todayDate, tomorrowDate time.Time, uid int) (*models.ProductDateItem, error) {
	// 获取 QuerySeter 对象，product 为表名
	qs := GetOrmer().QueryTable("product")

	var products []models.Product
	_, err := qs.RelatedSel("user").Filter("created__gte", todayDate).Filter("created__lt", tomorrowDate).
		OrderBy("-created").All(&products)

	if err != nil {
		return nil, err
	}
	item := models.ProductDateItem{
		Date:     utils.DateFormat(todayDate),
		Products: products,
		Uid:      uid,
	}
	return &item, nil
}
