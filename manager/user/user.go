package user

import (
	"productgo/dao"
	"productgo/models"
)

var (
	// Mgr 全局的一个 user manager
	Mgr = New()
)

type Manager interface {
	// Create 新建用户
	Create(username, password string) (int64, error)
	// Get 根据用户名和密码获取用户
	Get(username, password string) (*models.User, error)
	// IsExists 判断用户名是否已经存在
	IsExists(username string) bool
}

// New 返回 Manager 的默认实现
func New() Manager {
	return &manager{}
}

type manager struct{}

func (m *manager) Create(username, password string) (int64, error) {
	user := models.User{}
	user.Username = username
	user.Password = password
	return dao.AddUser(&user)
}

func (m *manager) Get(username, password string) (*models.User, error) {
	return dao.GetUserByParam(username, password)
}

func (m *manager) IsExists(username string) bool {
	return dao.UserExists(username)
}
