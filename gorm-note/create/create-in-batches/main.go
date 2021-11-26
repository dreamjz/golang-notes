package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	timeFormat = "2006-01-02 15:04:05.000"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %d, Birthday: %s\nCreatedAy: %s, UpadtedAt: %s, DeletedAt: %s",
		u.ID, u.Name, u.Age, u.Birthday.Format(timeFormat), u.CreatedAt.Format(timeFormat), u.UpdatedAt.Format(timeFormat),
		u.DeletedAt.Time.Format(timeFormat))
}

func main() {
	db, err := gorm.Open(sqlite.Open("create-in-batches.db"), &gorm.Config{
		Logger:          logger.Default.LogMode(logger.Info),
		CreateBatchSize: 3,
	})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	db.AutoMigrate(&User{})

	var users = []User{{Name: "user_1"}, {Name: "user_2"}, {Name: "user_3"}}
	result := db.Create(&users)
	log.Printf("Rows affected: %d,Error: %v", result.RowsAffected, result.Error)
	for _, user := range users {
		log.Printf("Inserted ID: %d", user.ID)
	}

	var users1 = []User{{Name: "user_1"}, {Name: "user_2"}, {Name: "user_3"}, {Name: "user_4"}, {Name: "user_5"}, {Name: "user_6"}}
	result = db.CreateInBatches(users1, 2)
	log.Printf("Rows affected: %d,Error: %v", result.RowsAffected, result.Error)
	for _, user := range users1 {
		log.Printf("Inserted ID: %d", user.ID)
	}

}
