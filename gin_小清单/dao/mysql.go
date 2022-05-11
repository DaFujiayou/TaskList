package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func Initmysql() (err error) {
	dsn := "root:123@tcp(127.0.0.1:3306)/bubble"
	DB, err = gorm.Open("mysql", dsn)
	return
}
