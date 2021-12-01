package main

import (
	"log"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DBName    = "update.db"
	UserCount = 10
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Group string
	Age   int
	Email Email
}

type Email struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Email  string
}

func main() {
	db, err := gorm.Open(sqlite.Open(DBName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("connect db failed: ", err.Error())
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Email{})
	CreateUsers(db, UserCount)

	// Update with SQL expression
	// UPDATE users SET age = age * 2 + 100 WHERE id = 1;
	db.Model(&User{}).Where("id = ?", 1).Update("age", gorm.Expr("age * ? + ?", 2, 100))
	// UPDATE users SET name = name || '123' WHERE id = 2 ;
	db.Model(&User{}).Where("id = ?", 2).Updates(map[string]interface{}{"name": gorm.Expr("name || ?", "_123")})
	// UPDATE users SET age = age * 100 WHERE id = 3 ;
	db.Model(&User{}).Where("id = ?", 3).UpdateColumn("age", gorm.Expr("age * ?", 100))
}

func CreateUsers(db *gorm.DB, num int) {
	var count int64
	db.Model(&User{}).Count(&count)
	if count > 0 {
		return
	}
	users := make([]User, num)
	for i := 0; i < num; i++ {
		grp := "group_" + strconv.Itoa(i%3)
		name := "user_" + strconv.Itoa(i)
		age := 10 + i
		email := name + "@example.com"
		users[i] = User{Name: name, Age: age, Group: grp, Email: Email{Email: email}}
	}
	db.Create(&users)
}
