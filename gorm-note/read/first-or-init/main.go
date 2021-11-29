package main

import (
	"gorm-note/utils"
	"log"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DBName    = "firstOrInit.db"
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

	// FirstOrInit
	// SELECT * FROM users WHERE name = 'non_existing' ORDER BY id LIMIT 1 ;
	var user User
	db.FirstOrInit(&user, User{Name: "non_existing"})
	utils.PrintRecord(user)
	// SELECT * FROM users WHERE name = 'user_0' ORDER BY id LIMIT 1 ;
	var user2 User
	db.FirstOrInit(&user2, map[string]interface{}{"name": "user_0"})
	utils.PrintRecord(user2)
	// SELECT * FROM users WHERE name = 'non_existing' ORDER BY id LIMIT 1 ;
	var user3 User
	db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrInit(&user3)
	utils.PrintRecord(user3)
	// SELECT * FROM users WHERE name = 'user_0' ORDER BY id LIMIT 1 ;	var user4 User
	var user4 User
	db.Where(User{Name: "user_0"}).Assign(User{Age: 20}).FirstOrInit(&user4)
	utils.PrintRecord(user4)
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
