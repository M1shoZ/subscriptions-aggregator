package models

import (
	"time"

	"github.com/M1shoZ/subscriptions-aggregator/utils"
	"github.com/google/uuid"
)

type MonthYear time.Time

type Subscriptions struct {
	// gorm.model
	ID          uint             `json:"id" gorm:"primary key;autoIncrement"`
	ServiceName string           `json:"service_name"`
	Price       uint             `json:"price"`
	UserID      uuid.UUID        `json:"user_id" gorm:"type:uuid"`
	StartDate   utils.MonthYear  `json:"start_date"`
	EndDate     *utils.MonthYear `json:"end_date,omitempty"` //опционально
}

type User struct {
	ID            uuid.UUID       `json:"id" gorm:"primary key; type:uuid;default:uuid_generate_v4()"`
	Subscriptions []Subscriptions `gorm:"foreignKey:UserID"`
}
