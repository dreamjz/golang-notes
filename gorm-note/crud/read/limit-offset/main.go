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
	DBName    = "limit.db"
	UserCount = 10
)

var (
	UserAllFields = []string{"id", "name", "age"}
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
		log.Fatal("connect db failed: ", err.Error())
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	db.AutoMigrate(&User{})
	CreateUsers(db, UserCount)

	// Limit&Offset
	// SELECT id,name,age FROM users LIMIT 3 ;
	var users []User
	db.Select(UserAllFields).Limit(3).Find(&users)
	utils.PrintRecord(users)
	// SELECT id,name,age FROM users LIMIT 10 ;
	// SELECT id,name,age FROM users ;
	var users1, users2 []User
	db.Select(UserAllFields).Limit(5).Find(&users1).Limit(-1).Find(&users2)
	utils.PrintRecord(users1, users2)
	// SELECT id,name,age FROM users OFFSET 3 ;
	var users3 []User
	db.Select(UserAllFields).Offset(3).Find(&users3)
	utils.PrintRecord(users3)
	// SELECT id,name,age FROM users OFFSET 3 LIMIT 4 ;
	var users4 []User
	db.Select(UserAllFields).Offset(3).Limit(4).Find(&users4)
	utils.PrintRecord(users4)
	// SELECT id,name,age FROM users OFFSET 3;
	// SELECT id,name,age FROM users ;
	var users5, users6 []User
	db.Select(UserAllFields).Offset(3).Find(&users5).Offset(-1).Find(&users6)
	utils.PrintRecord(users5, users6)

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
