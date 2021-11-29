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
	DBName    = "hooks.db"
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

func (u *User) AfterFind(*gorm.DB) (err error) {
	u.Age += 10
	return
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

	// Pluck
	// SELECT age FROM users ;
	var ages []int
	db.Model(&User{}).Pluck("age", &ages)
	utils.PrintRecord(ages)
	// SELECT age FROM users ;
	var ages2 []int
	db.Model(&User{}).Select("age").Find(&ages2)
	utils.PrintRecord(ages2)
	// SELECT name FROM users ;
	var names []string
	db.Model(&User{}).Pluck("name", &names)
	utils.PrintRecord(names)
	// SELECT DISTINCT `group` FROM users ;
	var groups []string
	db.Model(&User{}).Distinct().Pluck("group", &groups)
	utils.PrintRecord(groups)
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
