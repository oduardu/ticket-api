package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	FullName string
	Phone    int64
	Mail     string
	Category string
	Message  string
}

type TicketResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	FullName  string    `json:"fullname"`
	Phone     int64     `json:"phone"`
	Mail      string    `json:"mail"`
	Category  string    `json:"category"`
	Message   string    `json:"message"`
}
