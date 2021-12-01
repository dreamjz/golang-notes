package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DBName = "has-one.db"
)

type User struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Age        int
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
