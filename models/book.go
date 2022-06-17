package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	MediumPrice float32        `json:"medium_price"`
	Author      string         `json:"author"`
	ImageURL    string         `json:"img_url"`
	CreatedAt   time.Time      `json:"created"` //campo obrigatório do gorm
	UpdatedAt   time.Time      `json:"updated"` //campo obrigatório do gorm
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"` //campo obrigatório do gorm
}
