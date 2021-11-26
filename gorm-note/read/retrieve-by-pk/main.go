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
	db, err := gorm.Open(sqlite.Open("retrieve-by-pk.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect db: ", err.Error())
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.AutoMigrate(&User{})
	CreateUsers(db)

	// SELECT * FROM users WHERE id = 10 ORDER BY id LIMIT 1;
	var user User
	db.First(&user, 10)
	log.Printf("Record: %+v", user)
	// SELECT * FROM users WHERE id = 10 ORDER BY id LIMIT 1;
	var user2 User
	db.First(&user2, "10")
	log.Printf("Record: %+v", user2)
	// SELECT * FROM users WHERE id IN (1,2,3)
	var users []User
	db.Find(&users, []int{1, 2, 3})
	log.Printf("Records: %+v", users)
	// SELECT * FROM users;
	var users1 []User
	db.Find(&users1)
	log.Printf("Records: %+v", users1)
}

func CreateUsers(db *gorm.DB) {
	var count int64
	db.Model(&User{}).Count(&count)
	if count > 0 {
		return
	}
	users := make([]User, 10)
	for i := 1; i < 11; i++ {
		name := "user_" + strconv.Itoa(i)
		age := 10 + i
		users[i-1] = User{Name: name, Age: age}
	}
	db.Create(&users)
}
