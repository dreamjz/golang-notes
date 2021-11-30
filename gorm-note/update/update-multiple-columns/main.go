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

	// Update multiple columns
	// Update with struct
	// UPDATE users SET name = 'user_0_edited' WHERE id = 1;
	db.Model(&User{}).Where("id = ?", 1).Updates(User{Name: "user_0_edited", Age: 0})
	// Update with map
	// UPDATE users SET name = 'user_1_edited',age = 0 WHERE id = 2 ;
	user := User{ID: 2}
	db.Model(&user).Updates(map[string]interface{}{"name": "user_1_edited", "age": 0})
	utils.PrintRecord(user)
	// Update specified fields
	// UPDATE users SET name = 'user_2_edited',age = 0 WHERE id = 3 ;
	db.Model(&User{}).Where("id = ?", 3).Select("name", "age").Updates(User{Name: "user_2_edited"})
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
