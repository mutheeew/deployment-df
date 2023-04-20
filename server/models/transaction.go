package models

import "time"

type Transaction struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	User      User      `json:"user"`
	UserID    int       `json:"userID"`
	Attach    string    `json:"attach" gorm:"type:varchar(255)"`
	Status    string    `json:"status" gorm:"type: varchar(255)"`
	Price     int       `json:"price"`
}
