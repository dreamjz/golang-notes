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
	db, err := gorm.Open(sqlite.Open("single-obj.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect db: ", err.Error())
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.AutoMigrate(&User{})
	CreateUsers(db)

	var user, user1, user2 User
	db.First(&user)
	db.Take(&user1)
	db.Last(&user2)

	log.Printf("First user: %+v", user)
	log.Printf("One  of users: %+v", user1)
	log.Printf("Last user: %+v", user2)

	var user3 User
	// SELECT * FROM users ORDER BY id LIMIT 1;
	db.First(&user3)
	// SELECT * FROM users ORDER BY id LIMIT 1;
	ret := map[string]interface{}{}
	db.Model(&User{}).First(&ret)
	// invalid
	ret1 := map[string]interface{}{}
	db.Table("users").First(&ret1)
	// SELECT * FROM users LIMIT 1;
	ret2 := map[string]interface{}{}
	db.Table("users").Take(&ret2)
	// SELECT * FROM languages ORDER BY code LIMIT 1;
	type Language struct {
		Code string
		Name string
	}
	db.First(&Language{})
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
