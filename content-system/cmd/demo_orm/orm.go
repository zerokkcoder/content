package main

// import (
// 	"fmt"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type Account struct {
// 	ID        int64     `gorm:"column:id;primaryKey"`
// 	Username  string    `gorm:"column:username"`
// 	Password  string    `gorm:"column:password"`
// 	Nickname  string    `gorm:"column:nickname"`
// 	CreatedAt time.Time `gorm:"column:created_at"`
// 	UpdatedAt time.Time `gorm:"column:updated_at"`
// }

// func (*Account) TableName() string {
// 	return "account"
// }

// func main() {
// 	db := connDB()
// 	var accounts []Account
// 	if err := db.Find(&accounts).Error; err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(accounts)
// }

// func connDB() *gorm.DB {
// 	mysqlDB, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/cms_account?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	db, err := mysqlDB.DB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	db.SetMaxOpenConns(4)
// 	db.SetMaxIdleConns(2)

// 	mysqlDB = mysqlDB.Debug()

// 	return mysqlDB
// }
