package main

import (
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

func main() {
	db, err := gorm.Open(sqlite.Open("create-map.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	db.AutoMigrate(&User{})

	data := map[string]interface{}{
		"Name": "kesa-map",
		"Age":  18,
	}
	result := db.Model(&User{}).Create(data)
	log.Printf("Error %v", result.Error)
	log.Printf("Rows affected: %d", result.RowsAffected)

	datas := []map[string]interface{}{
		{"Name": "kesa_map_1", "Age": 10},
		{"Name": "kesa_map_2", "Age": 11},
		{"Name": "kesa_map_3", "Age": 12},
	}
	result = db.Model(&User{}).Create(datas)
	log.Printf("Error %v", result.Error)
	log.Printf("Rows affected: %d", result.RowsAffected)

}
