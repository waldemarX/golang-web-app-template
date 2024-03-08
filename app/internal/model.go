package model

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model

	ID       int `gorm:"primarykey"`
	Title    string
	SubTitle string
	Text     string
}
