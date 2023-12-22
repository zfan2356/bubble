package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func InitDatabase() (err error) {
	info := "root:123456@(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", info)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}
