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
	DBName    = "select.db"
	UserCount = 10
)

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open(DBName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("connected db failed:", err.Error())
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.AutoMigrate(&User{})
	CreateUsers(db, UserCount)

	// SELECT name FROM users ;
	var users []User
	db.Select("name").Find(&users)
	utils.PrintRecord(users)
	// SELECT name,age FROM users ;
	var users2 []User
	db.Select([]string{"name", "age"}).Find(&users2)
	utils.PrintRecord(users2)
}
func CreateUsers(db *gorm.DB, num int) {
	var count int64
	db.Model(&User{}).Count(&count)
	if count > 0 {
		return
	}
	users := make([]User, num)
	for i := 0; i < num; i++ {
		name := "user_" + strconv.Itoa(i)
		age := 10 + i
		users[i] = User{Name: name, Age: age}
	}
	db.Create(&users)
}
