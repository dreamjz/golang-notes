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

	// Update selected fields
	// UPDATE users SET name = 'user_0_new' WHERE id = 1 ;
	db.Model(&User{}).
		Where("id = ?", 1).
		Select("name").
		Updates(User{Name: "user_0_new"})
	// UPDATE users SET age = 10 WHERE id = 2 ;
	db.Table("users").Where("id = ?", 2).
		Omit("age").Updates(User{Name: "user_1_new", Age: 10})
	// Select all fields
	// UPDATE users SET id = 0,name = 'user_2_new',age = 0,group = '' WHERE id = 3;
	db.Table("users").Where("id = ?", 3).Select("*").
		Updates(User{Name: "user_2_new"})
	// Select all fields but omit name
	// UPDATE users SET id = 0,age = 0,group = '' WHERE id = 4;
	db.Table("users").Where("id = ?", 4).Select("*").
		Omit("name").Updates(User{Name: "user_3_new"})
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
