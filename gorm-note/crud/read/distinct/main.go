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
	DBName    = "distinct.db"
	UserCount = 10
)

var (
	GroupAndAvgAge = []string{"group", "AVG(age) as avg_age"}
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Group string
	Age   int
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
	CreateUsers(db, UserCount)

	// Distinct
	// SELECT DISTINCT name FROM users ;
	var result []map[string]interface{}
	db.Model(&User{}).Distinct("group").Find(&result)
	utils.PrintRecord(result)

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
		users[i] = User{Name: name, Age: age, Group: grp}
	}
	db.Create(&users)
}
