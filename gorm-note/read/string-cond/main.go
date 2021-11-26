package main

import (
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

func main() {
	db, err := gorm.Open(sqlite.Open("string-cond.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed connect to db: ", err.Error())
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.AutoMigrate(&User{})
	CreateUsers(db, 10)

	// String
	StringCondition(db)
}

func CreateUsers(db *gorm.DB, num int) {
	var count int64
	db.Model(&User{}).Count(&count)
	if count > 0 {
		return
	}
	users := make([]User, num)
	for i := 0; i < num; i++ {
		name := "user_" + strconv.Itoa(i+1)
		age := 10 + i
		users[i] = User{Name: name, Age: age}
	}
	db.Create(&users)
}

func PrintRecord(record interface{}) {
	switch val := record.(type) {
	case User:
		log.Printf("Record: %+v", val)
	case []User:
		log.Printf("Records: %+v", val)
	}
}

func StringCondition(db *gorm.DB) {
	// SELECT * FROM users WHERE name = 'user' ORDER BY id LIMIT 1;
	var user User
	db.Where("name = ?", "user").First(&user)
	PrintRecord(user)
	// SELECT * FROM users WHERE name <> 'user';
	var users []User
	db.Where("name <> ?", "user").Find(&users)
	PrintRecord(users)
	// SELECT * FROM users WHERE name IN ('user_1','user_2','user_3')
	var users1 []User
	db.Where("name IN ?", []string{"user_1", "user_2", "user_3"}).Find(&users1)
	PrintRecord(users1)
	// SELECT * FROM users WHERE name LIKE '%user%';
	var users3 []User
	db.Where("name LIKE ?", "%user%").Find(&users3)
	PrintRecord(users3)
	// SELECT * FROM users WHERE name = 'user_1' AND age = 11;
	var users4 []User
	db.Where("name = ? AND age = ?", "user_1", 11).Find(&users4)
	PrintRecord(users4)
	// SELECT * FROM users WHERE age <= 10;
	var users5 []User
	db.Where("age <= ?", 15).Find(&users5)
	PrintRecord(users5)
	// SELECT * FROM users WHERE age BETWEEN 10 AND 16;
	var users6 []User
	db.Where("age BETWEEN ? AND ?", 10, 16).Find(&users6)
	PrintRecord(users6)
}
