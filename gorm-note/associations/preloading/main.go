package main

import (
	"gorm-note/utils"
	"log"
	"math/rand"
	"strconv"

	"gorm.io/gorm/clause"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DBName    = "preloading.db"
	UserCount = 10
)

type User struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Orders     []Order
	CreditCard CreditCard
}

type Order struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Price  float64
}

type CreditCard struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Number string
}

func main() {
	db := initializeDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	createTables(db)
	createUsers(db, UserCount)

	// Preloading
	// SELECT * FROM `orders` WHERE `orders`.`user_id` = 1
	// SELECT * FROM `users` WHERE id = 1
	var user User
	db.Preload("Orders").Where("id = ?", 1).Find(&user)
	utils.PrintRecord(user)
	// Joins preload work with has-one or belongs-to
	// SELECT `users`.`id`,`users`.`name`,`CreditCard`.`id` AS `CreditCard__id`,`CreditCard`.`user_id` AS `CreditCard__user_id`,`CreditCard`.`number` AS `CreditCard__number` FROM `users`
	// LEFT JOIN `credit_cards` `CreditCard` ON `users`.`id` = `CreditCard`.`user_id` WHERE `users`.`id` = 2 LIMIT 1
	var user3 User
	db.Joins("CreditCard").Take(&user3, 2)
	utils.PrintRecord(user3)
	// not working, will cause panic
	joinsPreload(db)
	// SELECT * FROM `credit_cards` WHERE `credit_cards`.`user_id` IN (1,2,3,4,5,6,7,8,9,10)
	// SELECT * FROM `orders` WHERE `orders`.`user_id` IN (1,2,3,4,5,6,7,8,9,10)
	// SELECT * FROM `users`
	var users []User
	db.Preload(clause.Associations).Find(&users)
	utils.PrintRecord(users)
	// preload with conditions
	// SELECT * FROM `orders` WHERE `orders`.`user_id` = 2 AND price > 0.500000
	// SELECT * FROM `users` WHERE id = 2
	var user4 User
	db.Where("id = ?", 2).Preload("Orders", "price > ?", 0.5).Find(&user4)
	utils.PrintRecord(user4)
	// custom preloading SQL
	// SELECT * FROM `orders` WHERE `orders`.`user_id` = 2 ORDER BY price
	// SELECT * FROM `users` WHERE id = 2
	var user5 User
	db.Where(&User{ID: 2}).Preload("Orders", func(db *gorm.DB) *gorm.DB {
		return db.Order("price")
	}).Find(&user5)
	utils.PrintRecord(user5)

}

func joinsPreload(db *gorm.DB) {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("Panic: %v", e)
		}
	}()
	db.Joins("Orders").Find(&User{}, 2)
}

func initializeDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(DBName), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		QueryFields: true,
	})
	if err != nil {
		log.Fatalln("connect db failed: ", err.Error())
	}
	return db
}

func createTables(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Order{})
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
		price := rand.Float64() * float64(i)
		price2 := rand.Float64() * float64(i)
		creditCardNum := "xxx_" + strconv.Itoa(i)
		users[i] = User{Name: name, Orders: []Order{{Price: price}, {Price: price2}}, CreditCard: CreditCard{Number: creditCardNum}}
	}
	db.Create(&users)
}
