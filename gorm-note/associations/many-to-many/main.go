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
	DBName    = "many-to-many.db"
	UserCount = 10
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func main() {
	db := initializeDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	createTables(db)
	createUsers(db, UserCount)

	// SELECT `languages`.`id`,`languages`.`name` FROM `languages`
	// JOIN `user_languages` ON `user_languages`.`language_id` = `languages`.`id`
	// AND `user_languages`.`user_id` = 1
	// WHERE name IN ('lang_0')
	var languages []Language
	db.Model(&User{ID: 1}).Where("name IN ?", []string{"lang_0"}).Association("Languages").Find(&languages)
	utils.PrintRecord(languages)
	// append new association
	db.Model(&User{ID: 1}).Association("Languages").Append(&Language{Name: "new_lang"})
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
	db.AutoMigrate(&Language{})
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
		lang := "lang_" + strconv.Itoa(i)
		lang2 := "lang_" + strconv.Itoa(i+1)
		users[i] = User{Name: name, Languages: []Language{{Name: lang}, {Name: lang2}}}
	}
	db.Create(&users)
}
