package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email string `json:"email" gorm:"unique;not null;"`
	Name  string `json:"name" gorm:"not null"`

	Desc string `json:"desc"`
}
