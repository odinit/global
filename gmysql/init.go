/*
初始化MySQL连接
*/

package gmysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(host string, port int, user, password, dbname string, opts ...gorm.Option) (err error) {
	DB, err = gorm.Open(
		mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)),
		opts...,
	)
	return
}

func DefaultInit() (err error) {
	return Init("localhost", 3306, "root", "123456", "test")
}
