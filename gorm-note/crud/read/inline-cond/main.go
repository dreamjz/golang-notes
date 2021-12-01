package main

import (
	"gorm-note/utils"
	"log"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Age  int
}

const (
	RecordsNum = 10
	DBName     = "condition.db"
)

func main() {
	db, err := gorm.Open(sqlite.Open(DBName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("connect to db failed: ", err.Error())
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.AutoMigrate(&User{})
	CreateUsers(db, RecordsNum)

	// Inline Conditions
	// SELECT * FROM users WHERE id = 1 ORDER BY id LIMIT 1;
	var user User
	db.First(&user, "id = ?", 1)
	utils.PrintRecord(user)
	// SELECT * FROM users WHERE name = 'user_1' LIMIT 1;
	var user1 User
	db.Take(&user1, "name = ?", "user_1")
	utils.PrintRecord(user1)
	// SELECT * FROM users WHERE name <> 'user_2' AND age >= 15;
	var users []User
	db.Find(&users, "name <> ? AND age >= ?", "user_2", 15)
	utils.PrintRecord(users)
	// SELECT * FROM users WHERE age = 16 LIMIT 1;
	var user2 User
	db.Take(&user2, User{Age: 16})
	utils.PrintRecord(user2)
	// SELECT * FROM users WHERE age = 13 LIMIT 1;
	var user3 User
	db.Take(&user3, map[string]interface{}{"Age": 13})
	utils.PrintRecord(user3)
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
