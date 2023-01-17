package db

import (
	"fmt"
	"thegame/model"

	"thegame/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config config.DBConfig) (*gorm.DB, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@database:%s/%s", config.Username, config.Password, config.DbPort, config.Database)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.GameState{})

	fmt.Println("Connected to DB ...")
	return db, nil
}
