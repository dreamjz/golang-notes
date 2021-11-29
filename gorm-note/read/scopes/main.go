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
	DBName    = "scopes.db"
	UserCount = 10
)

var (
	UserAllFields = []string{"id", "name", "age", "group"}
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

	// Scopes
	// SELECT id,name,age,group FROM users WHERE age > 15 ;
	var users []User
	db.Select(UserAllFields).
		Scopes(AgeGreaterThan15).
		Find(&users)
	utils.PrintRecord(users)
	// SELECT id,name,age,group FROM users WHERE group IN ('group_0','group_2') ;
	var users2 []User
	db.Select(UserAllFields).
		Scopes(GroupIn([]string{"group_0", "group_2"})).
		Find(&users2)
	utils.PrintRecord(users2)
}

func AgeGreaterThan15(db *gorm.DB) *gorm.DB {
	return db.Where("age > ?", 15)
}

func GroupIn(groups []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("`group` IN (?)", groups)
	}
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
