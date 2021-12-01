package main

import (
	"database/sql"
	"gorm-note/utils"
	"log"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DBName    = "advanced-query.db"
	UserCount = 10
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Group string
	Age   int
	Email Email
}

type Email struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Email  string
}

func main() {
	db, err := gorm.Open(sqlite.Open(DBName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("connect db failed: ", err.Error())
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Email{})
	CreateUsers(db, UserCount)

	// SubQuery
	// SELECT name,age FROM users WHERE age > (SELECT AVG(age) FROM users) ;
	var results []map[string]interface{}
	subQuery := db.Model(&User{}).Select("AVG(age)")
	db.Model(&User{}).Select("name", "age").Where("age > (?)", subQuery).Find(&results)
	utils.PrintRecord(results)
	// SELECT * FROM (SELECT name,age FROM users) as u WHERE age < 15 ;
	var results1 []map[string]interface{}
	db.Table("(?) as u", db.Model(&User{}).Select("name", "age")).Where("age < ?", 15).Find(&results1)
	utils.PrintRecord(results1)
	// SELECT * FROM users WHERE (name = 'user_0' AND age = 10) OR (age > 15) ;
	var results2 []map[string]interface{}
	db.Model(&User{}).Where(
		db.Where("name = ?", "user_0").Where("age = ?", 10),
	).Or(
		db.Where("age > 15"),
	).Find(&results2)
	utils.PrintRecord(results2)
	// SELECT * FROM users WHERE (name,age) IN (('user_0', 10),('user_1', 11),('user_2', 12));
	var results3 []map[string]interface{}
	db.Select("name", "age").
		Model(&User{}).
		Where("(name,age) IN ?", [][]interface{}{{"user_0", 10}, {"user_1", 11}, {"user_2", 12}}).
		Find(results3)
	utils.PrintRecord(results3)
	// SELECT * FROM users WHERE name = 'user_1' OR age = 16;
	var results4 []map[string]interface{}
	db.Model(&User{}).
		Select("name", "age").
		Where("name = @name OR age = @age", sql.Named("name", "user_0"), sql.Named("age", 16)).
		Find(&results4)
	utils.PrintRecord(results4)
	// SELECT * FROM users WHERE name = 'user_1' OR age = 16;
	var results5 []map[string]interface{}
	db.Model(&User{}).
		Select("name", "age").
		Where("name = @name OR age = @age", map[string]interface{}{"name": "user_1", "age": 16}).
		Find(&results5)
	utils.PrintRecord(results5)

}

func CreateUsers(db *gorm.DB, num int) {
	var count int64
	db.Model(&User{}).Count(&count)
	if count > 0 {
		return
	}
	users := make([]User, num)
	for i := 0; i < num; i++ {
		grp := "group_" + strconv.Itoa(i%3)
		name := "user_" + strconv.Itoa(i)
		age := 10 + i
		email := name + "@example.com"
		users[i] = User{Name: name, Age: age, Group: grp, Email: Email{Email: email}}
	}
	db.Create(&users)
}
