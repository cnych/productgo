package utils

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() error {
	db0, err := orm.GetDB()
	if err != nil {
		return err
	}
	db = db0
	return nil
}
//func InitMySQL() error {
//	if db == nil {
//		db0, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/productgo")
//		if err != nil {
//			return err
//		}
//		db = db0
//		// 初始化User表
//		err = InitTableUser()
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

func InitTableUser() error {
	sql := `CREATE TABLE IF NOT EXISTS user(
				id INT(10) PRIMARY KEY AUTO_INCREMENT NOT NULL,
				username VARCHAR(64),
				password VARCHAR(64),
				createtime int(10)
			);`
	if _, err := ExecSQL(sql); err != nil {
		return err
	}
	return nil
}

func ExecSQL(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func QueryRow(sql string) *sql.Row {
	return db.QueryRow(sql)
}
