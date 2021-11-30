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

	// Update single column
	// Update with conditions
	// UPDATE users SET name = 'user_0_edited' WHERE name = 'user_0' ;
	db.Model(&User{}).Where("name = ?", "user_0").Update("name", "user_0_edited")
	// UPDATE users SET name = 'user_1_edited' WHERE id = 2 ;
	user := User{ID: 2}
	db.Model(&user).Update("name", "user_1_edited")
	utils.PrintRecord(user)
	// Update with conditions and model value
	// UPDATE users SET name = 'user_2_edited' WHERE id = 3 AND group = 'group_2' AND name = 'user_2'
	user2 := User{ID: 3, Group: "group_2"}
	db.Model(&user2).Where("name = ?", "user_2").Update("name", "user_2_edited")
	utils.PrintRecord(user2)
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
