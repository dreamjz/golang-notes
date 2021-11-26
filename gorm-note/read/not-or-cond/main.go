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

	// Not Conditions
	// SELECT * FROM users WHERE NOT name = 'user_1' LIMIT 1;
	var users []User
	db.Not("name = ?", "user_1").Find(&users)
	utils.PrintRecord(users)
	// SELECT * FROM users WHERE name NOT IN ('user_2','user_3','user_4');
	var users1 []User
	db.Not(map[string]interface{}{"name": []string{"user_2", "user_3", "user_4"}}).Find(&users1)
	utils.PrintRecord(users1)
	// SELECT * FROM users WHERE name <> 'user_5' AND age <> 11;
	var users2 []User
	db.Not(User{Name: "user_5", Age: 11}).Find(&users2)
	utils.PrintRecord(users2)
	// SELECT * FROM users WHERE id NOT IN (1,2,3);
	var users3 []User
	db.Not([]int64{1, 2, 3}).Find(&users3)
	utils.PrintRecord(users3)

	// Or Conditions
	// SELECT * FROM users WHERE name = 'user_1' OR age = 16;
	var users4 []User
	db.Where("name = ?", "user_1").Or("age = ?", 16).Find(&users4)
	utils.PrintRecord(users4)
	// SELECT * FROM user WHERE name = 'user_2' OR (name = 'user_3' AND age = 13) ;
	var users5 []User
	db.Where("name = 'user_2'").Or(User{Name: "user_3", Age: 13}).Find(&users5)
	utils.PrintRecord(users5)
	// SELECT * FROM users WHERE age = 10 OR (name = 'user_0' AND age = 10);
	var users6 []User
	db.Where("age = ?", 10).Or(map[string]interface{}{"name": "user_0", "age": 10}).Find(&users6)
	utils.PrintRecord(users6)
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
