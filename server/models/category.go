package models

type Category struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" binding:"required, name" gorm:"unique; not null"`
}
