package utils

import (
	"fmt"
	"lamb-code/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDB(dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USER,
		config.DB_PW,
		config.DB_ADDR,
		dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}
	return db, nil
}
