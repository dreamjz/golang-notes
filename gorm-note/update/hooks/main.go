package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

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

func (User) BeforeUpdate(db *gorm.DB) (err error) {
	time.Sleep(1 * time.Second)
	fmt.Println("Before update")
	return
}

func (User) BeforeSave(db *gorm.DB) (err error) {
	time.Sleep(1 * time.Second)
	fmt.Println("Before save")
	return
}

func (User) AfterUpdate(db *gorm.DB) (err error) {
	time.Sleep(1 * time.Second)
	fmt.Println("After update")
	return
}

func (User) AfterSave(db *gorm.DB) (err error) {
	time.Sleep(1 * time.Second)
	fmt.Println("After save")
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

	// Update hooks
	db.Model(&User{}).Where("id = ?", 1).Update("name", "user_0_new")
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
