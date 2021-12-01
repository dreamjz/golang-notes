package main

import (
	"gorm-note/utils"
	"log"
	"strconv"

	"gorm.io/plugin/soft_delete"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DBName    = "delete.db"
	UserCount = 10
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Group     string
	Age       int
	DeletedAt soft_delete.DeletedAt
	Email     Email
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

	// Raw SQL
	// Scan
	var result map[string]interface{}
	db.Raw("SELECT id,name,age FROM users WHERE id = ?", 1).Scan(&result)
	utils.PrintRecord(result)
	var results []map[string]interface{}
	db.Raw("SELECT id,name FROM users WHERE age > ?", 10).Scan(&results)
	utils.PrintRecord(results)
	var sumAge int
	db.Raw("SELECT SUM(age) FROM users").Scan(&sumAge)
	log.Println("Sum age: ", sumAge)
	// Exec
	db.Exec("UPDATE users SET age = age + ?", 10)
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
