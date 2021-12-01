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
	DBName    = "to-sql.db"
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

	// Row and Rows
	// Row
	var (
		name string
		age  int
	)
	row := db.Model(&User{}).Where("age > 10", "user_0").Order("age desc").
		Select("name", "age").Row()
	row.Scan(&name, &age)
	log.Printf("Name: %s,Age: %d", name, age)
	// Rows
	rows, _ := db.Model(&User{}).Where("age > ?", 15).Select("name", "age").Rows()
	defer rows.Close()
	for rows.Next() {
		var name string
		var age int
		rows.Scan(&name, &age)
		log.Printf("Row[name: %s,Age: %d]", name, age)
		// Scan rows to struct
		var user User
		db.ScanRows(rows, &user)
		utils.PrintRecord(user)
	}

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
