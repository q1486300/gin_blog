package model

import (
	"fmt"
	"gin_blog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB

func InitDb() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// gorm 日誌模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外鍵約束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用預設 Transaction (提高執行速度)
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 單數表名
		},
	})
	if err != nil {
		panic("連接資料庫失敗，請檢查參數: " + err.Error())
	}

	db.AutoMigrate(&User{}, &Article{}, &Category{})

	sqlDb, _ := db.DB()

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(10 * time.Second)
}
