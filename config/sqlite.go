package config

import (
	"fmt"
	"os"

	"github.com/oduardu/ticket-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		fmt.Printf("[ INFO ] - Database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		fmt.Errorf("SQLite Ticket Error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Ticket{})
	if err != nil {
		fmt.Errorf("SQLite Automigration Error: %v", err)
		return nil, err
	}

	return db, nil
}
