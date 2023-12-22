package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	if err := dao.InitDatabase(); err != nil {
		panic(err)
	}
	defer dao.DB.Close()

	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run(":8080")
}
