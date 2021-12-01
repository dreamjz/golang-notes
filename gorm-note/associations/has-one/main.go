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
	DBName    = "has-one.db"
	UserCount = 10
)

type User struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	CreditCard CreditCard
}

type CreditCard struct {
	ID     uint `gorm:"primaryKey"`
	Number string
	UserID uint
}

func main() {
	db := initializeDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	createTables(db)
	createUsers(db, UserCount)

	//query association
	var creditCard CreditCard
	db.Model(&User{ID: 1}).Association("CreditCard").Find(&creditCard)
	utils.PrintRecord(creditCard)
	// append association
	// SELECT * FROM `credit_cards` WHERE `credit_cards`.`user_id` = 1
	// INSERT INTO `credit_cards` (`number`,`user_id`) VALUES ("new_xxx",2) ON CONFLICT (`id`) DO UPDATE SET `user_id`=`excluded`.`user_id` RETURNING `id`
	db.Model(&User{ID: 2}).Association("CreditCard").Append(&CreditCard{Number: "new_xxx"})
}

func initializeDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DBName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("connect db failed: ", err.Error())
	}
	return db
}

func createTables(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&CreditCard{})
}

func createUsers(db *gorm.DB, num int) {
	var count int64
	db.Model(&User{}).Count(&count)
	if count > 0 {
		return
	}
	users := make([]User, num)
	for i := 0; i < num; i++ {
		name := "user_" + strconv.Itoa(i)
		creditCardNumber := "xxx" + strconv.Itoa(i)
		users[i] = User{Name: name, CreditCard: CreditCard{Number: creditCardNumber}}
	}
	db.Create(&users)
}
