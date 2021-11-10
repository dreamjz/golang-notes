package core

import (
	"fmt"
	"go-gin-example/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Gorm initialize database connection
func Gorm() {
	// mysql config
	m := global.AppConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User, m.Password, m.Host, m.Port, m.Name)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.AppConfig.Database.TablePrefix, // set table name prefix
			SingularTable: true,                                  // use singular table name
		},
	})
	if err != nil {
		log.Fatalf("open database failed, %s", err)
	}
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.AppDB = db
}

func CloseDB() {
	sqlDB, _ := global.AppDB.DB()
	defer sqlDB.Close()
}
