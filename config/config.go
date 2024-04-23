package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("Error Initializing config: %v", err)
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}
