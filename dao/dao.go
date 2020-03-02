package dao

import (
	"sync"

	"github.com/astaxie/beego/orm"
)

var globalOrm orm.Ormer
var once sync.Once

func GetOrmer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
