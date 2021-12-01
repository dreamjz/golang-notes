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

	// Struct and Map
	StructMapCondition(db)
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

func StructMapCondition(db *gorm.DB) {
	// Struct
	// SELECT * FROM users WHERE name = 'user_1' AND age = 10 ORDER BY id LIMIT 1;
	var user User
	db.Where(&User{Name: "user_1", Age: 10}).First(&user)
	PrintRecord(user)
	// SELECT * FROM users WHERE name = 'user_2' AND age = 11;
	var user2 User
	cond := map[string]interface{}{"Name": "user_2", "Age": 11}
	db.Where(cond).Find(&user2)
	PrintRecord(user2)
	// SELECT * FROM users WHERE id IN (1,2,3)
	var users []User
	db.Where([]int{1, 2, 3}).Find(&users)
	PrintRecord(users)

	// SELECT * FROM users WHERE name = 'user_1' LIMIT 1;
	var user3 User
	db.Where(&User{Name: "user_1", Age: 0}).Take(&user3)
	PrintRecord(user3)
	// SELECT * FROM users WHERE name = 'user_2' AND age = 0 LIMIT 1;
	var user4 User
	db.Where(map[string]interface{}{"Name": "user_2", "Age": 0}).Take(&user4)
	PrintRecord(user4)
	// SELECT * FROM users WHERE name = 'user_3' AND age = 12 LIMIT 1;
	var user5 User
	db.Where(&User{Name: "user_3", Age: 12}, "name", "age").Take(&user5)
	PrintRecord(user5)
	// SELECT * FROM users WHERE age = 0 LIMIT 1;
	var user6 User
	db.Where(&User{Name: "user_3"}, "age").Take(&user6)
	PrintRecord(user6)
}
