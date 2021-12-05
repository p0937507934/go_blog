package db

import (
	"fmt"

	"github.com/blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		"admin",
		"l5341500",
		"database-1.c1v8bdx2y2xw.us-east-2.rds.amazonaws.com",
		"3306",
		"indochat",
		"charset=utf8&parseTime=true")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
}
