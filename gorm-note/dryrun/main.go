package main

import (
	"database/sql"
	"gorm-note/utils"
	"log"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DBName    = "raw.db"
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

	// Named argument
	// SELECT * FROM users WHERE name = 'user_0' LIMIT 1;
	var user User
	db.Where("name = @name", sql.Named("name", "user_0")).Take(&user)
	utils.PrintRecord(user)
	// SELECT * FROM users WHERE age = 12 LIMIT 1 ;
	var user2 User
	db.Where("age = @age ", map[string]interface{}{"age": 12}).Take(&user2)
	utils.PrintRecord(user2)
	// Named Argument with Raw SQL
	var result []map[string]interface{}
	db.Raw("SELECT id,name FROM users WHERE age = @age OR id = @id ", sql.Named("age", 15), sql.Named("id", 3)).
		Scan(&result)
	utils.PrintRecord(result)
	var result2 map[string]interface{}
	db.Raw("SELECT name,age FROM users WHERE id = @id OR age > @age", map[string]interface{}{"id": 2, "age": 13}).
		Scan(&result2)
	utils.PrintRecord(result2)
	// struct
	type NameArg struct {
		Name string
		Age  int
	}
	var result3 map[string]interface{}
	db.Raw("SELECT * FROM users WHERE name = @Name AND age = @Age", NameArg{Name: "user_0", Age: 10}).Scan(&result3)
	utils.PrintRecord(result3)
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
