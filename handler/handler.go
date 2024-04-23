package handler

import (
	"github.com/oduardu/ticket-api/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeHandler() {
	db = config.GetSQLite()
}
