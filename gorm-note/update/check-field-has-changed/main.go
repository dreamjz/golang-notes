package main

import (
	"fmt"
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

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// if age changed
	if tx.Statement.Changed("name") {
		fmt.Println("name changed")
		tx.Statement.SetColumn("age", 100)
		return
	}
	// if name or age changed
	if tx.Statement.Changed("name", "age") {
		fmt.Println("name or age changed")
		return
	}
	// if any fields changed
	if tx.Statement.Changed() {
		fmt.Println("any fields changed")
		return
	}
	return
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

	// Check Field has changed
	// UPDATE users SET name = 'new_name',age = 100  WHERE id = 1 ;
	db.Model(&User{ID: 1, Name: "user_0"}).Update("name", "new_name")
	// UPDATE users SET age = age * 2 WHERE id =2  ;
	db.Model(&User{ID: 2, Age: 11}).Update("age", gorm.Expr("age * ?", 2))
	// UPDATE users SET name = 'new_name',age = 1300 WHERE id = 3 ;
	db.Model(&User{ID: 3, Age: 12, Name: "user_2"}).Updates(User{Name: "new_name", Age: 1300})
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
