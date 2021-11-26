package main

import (
	"log"

	"gorm.io/gorm/clause"

	"gorm.io/gorm/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("upsert.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&CreditCard{})

	user := User{Name: "kesa", CreditCard: CreditCard{Number: "12346"}}
	result := db.Create(&user)
	log.Printf("New user ID: %d, CreditCard ID: %d", user.ID, user.CreditCard.ID)
	log.Printf("Rows affected: %d, Error: %v", result.RowsAffected, result.Error)

	user1 := User{Name: "kesa", CreditCard: CreditCard{Number: "78910"}}
	result = db.Omit(clause.Associations).Create(&user1)
	log.Printf("New user ID: %d, CreditCard ID: %d", user1.ID, user1.CreditCard.ID)
	log.Printf("Rows affected: %d, Error: %v", result.RowsAffected, result.Error)

}
