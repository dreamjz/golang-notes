package main

import (
	"log"

	"gorm.io/gorm/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("quick-start.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	db.AutoMigrate(&Product{})

	db.Create(&Product{
		Code:  "D42",
		Price: 100,
	})

	var product Product
	db.First(&product, 1)
	log.Printf("Result:%+v\n", product)
	db.First(&product, "code = ?", "D42")
	log.Printf("Result:%+v\n", product)

	db.Model(&product).Update("Price", 200)
	log.Printf("Result:%+v\n", product)
	db.Model(&product).Updates(Product{Code: "F42", Price: 300})
	log.Printf("Result:%+v\n", product)
	db.Model(&product).Updates(map[string]interface{}{"Code": "E42", "Price": 150})
	log.Printf("Result:%+v\n", product)

	db.Delete(&product, 1)
	log.Printf("Result:%+v\n", product)

}
