package product

import (
	"fmt"
	"productgo/dao"
	"productgo/models"
	"time"
)

var (
	// Mgr 全局的一个 product manager
	Mgr = New()
)

type Manager interface {
	// Create
	Create(name, description, link string, user *models.User) (int64, error)
	// Get
	Get(interface{}) (*models.Product, error)
	// Vote
	Vote(*models.User, *models.Product) error
	// GetProductsByDate
	ListByDate(todayDate, tomorrowDate time.Time, uid int) (*models.ProductDateItem, error)
}

// New 返回 Manager 的默认实现
func New() Manager {
	return &manager{}
}

type manager struct{}

func (m *manager) Create(name, description, link string, user *models.User) (int64, error) {
	product := models.Product{}
	product.Name = name
	product.Description = description
	product.Link = link
	product.User = user
	// 其他的一些操作
	return dao.AddProduct(&product)
}

func (m *manager) Get(idOrName interface{}) (*models.Product, error) {
	id, ok := idOrName.(int64)
	if ok {
		return dao.GetProductById(id)
	}
	return nil, fmt.Errorf("invalid parameter: %v, should be ID(int64)", idOrName)
	// todo，其他类型的查询
}

func (m *manager) Vote(user *models.User, product *models.Product) error {
	return dao.VoteProduct(user, product)
}

func (m *manager) ListByDate(todayDate, tomorrowDate time.Time, uid int) (*models.ProductDateItem, error) {
	return dao.GetProductsByDate(todayDate, tomorrowDate, uid)
}
