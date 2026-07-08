package store

import (
	"github.com/312022151125/go-backend-template/internal/conf"
	"github.com/312022151125/go-backend-template/internal/model"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() error {
	var err error
	gormConfig := gorm.Config{Logger: logger.Discard}
	if conf.IsDebug() {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	db, err = gorm.Open(sqlite.Open(conf.AppConfig.Database.Path+"?_journal_mode=WAL&_busy_timeout=5000&_foreign_keys=ON"), &gormConfig)
	if err != nil {
		return err
	}
	return db.AutoMigrate(new(model.User))
}

func Close() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
