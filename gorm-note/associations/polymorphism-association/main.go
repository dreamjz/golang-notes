package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DBName = "polymorphism-association.db"
)

type Cat struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Toys []Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	OwnerID   uint
	OwnerType string
}

func main() {
	db := initializeDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	createTables(db)

	// INSERT INTO cats (name) VALUES ('cat_1');
	// INSERT INTO toys (name,owner_id,owner_type) VALUES ('toy_1','1','cats'),('toy_2','1','cats'),('toy_3','1','cats');
	db.Create(&Cat{Name: "cat_1", Toys: []Toy{{Name: "toy_1"}, {Name: "toy_2"}, {Name: "toy_3"}}})
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
	db.AutoMigrate(&Cat{})
	db.AutoMigrate(&Dog{})
	db.AutoMigrate(&Toy{})
}
