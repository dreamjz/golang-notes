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
	DBName    = "group.db"
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

type GrpAvgAge struct {
	Group  string
	AvgAge float64
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

	// Group By & Having
	// SELECT group,AVG(age) as avg_age FROM users GROUP BY `group` ;
	var results []GrpAvgAge
	db.Model(&User{}).Select(GroupAndAvgAge).Group("group").Find(&results)
	utils.PrintRecord(results)
	// SELECT group,AVG(age) as avg_age FROM users GROUP BY `group` HAVING avg_age > 14 ;
	var results1 []GrpAvgAge
	db.Model(&User{}).Select(GroupAndAvgAge).Group("group").Having("avg_age > ?", 14).Find(&results1)
	utils.PrintRecord(results1)
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
