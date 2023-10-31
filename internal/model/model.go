package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"yuwang.idv/go-tour/blog-api/global"
	"yuwang.idv/go-tour/blog-api/pkg/setting"
)

// Common fields
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	IsDel      uint8  `json:"is_del"`
}

// https://github.com/go-gorm/postgres
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	var dbLogger logger.Interface

	if global.ServerSetting.RunMode == "debug" {
		dbLogger = logger.Default.LogMode(logger.Info)
	}
	conStr := fmt.Sprintf(
		//"host=your_postgres_host user=your_user password=your_password dbname=your_db port=5432 sslmode=disable"
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		databaseSetting.Host,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.DBName,
		databaseSetting.Port,
		"Asia/Shanghai",
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  conStr,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
