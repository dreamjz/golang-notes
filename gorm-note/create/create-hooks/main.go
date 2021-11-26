package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

const (
	timeFormat = "2006-01-02 15:04:05.000"
)

func (u User) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Age: %d, Birthday: %s\nCreatedAy: %s, UpadtedAt: %s, DeletedAt: %s",
		u.ID, u.Name, u.Age, u.Birthday.Format(timeFormat), u.CreatedAt.Format(timeFormat), u.UpdatedAt.Format(timeFormat),
		u.DeletedAt.Time.Format(timeFormat))
}

func (User) BeforeCreate(db *gorm.DB) error {
	log.Println("Before Create")
	return nil
}

func (User) BeforeSave(db *gorm.DB) error {
	log.Println("Before Save")
	return nil
}

func (User) AfterCreate(db *gorm.DB) error {
	log.Println("After Create")
	return nil
}

func (u User) AfterSave(db *gorm.DB) error {
	log.Println("After Save")
	if u.Age < 21 {
		return errors.New("illegal age,roll back")
	}
	return nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("create-hooks.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	db.AutoMigrate(&User{})

	user := User{Name: "kesa", Age: 18, Birthday: time.Now()}
	result := db.Create(&user)
	log.Printf("Rows affected: %d, Error: %v", result.RowsAffected, result.Error)
	log.Printf("New record ID: %d", user.ID)

}
