package db

import (
	"task-api/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DBName = "tasks.db"

type DB struct {
	DB *gorm.DB
}

func New() *DB {
	db, err := gorm.Open(sqlite.Open(DBName), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Task{})

	return &DB{
		DB: db,
	}
}
