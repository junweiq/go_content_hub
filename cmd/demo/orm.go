package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Nickname  string    `gorm:"column:nickname"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:created_at"`
}

func (a Account) TableName() string {
	return "account"
}

func main() {
	db := connDB()
	//var accounts []Account
	//if err := db.Find(&accounts).Error; err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(accounts)

	var account Account
	if err := db.Where("id = ?", 1).First(&account).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}

func connDB() *gorm.DB {
	mysqlDB, err := gorm.Open(mysql.Open("root:handsome@tcp(localhost:3306)/cms_user?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(4) // 最大連接數
	db.SetMaxIdleConns(2) // 通常/2

	//if env == "test" {
	mysqlDB = mysqlDB.Debug()
	//}

	return mysqlDB
}
