package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model

	Number   string `json:"number"`
	Capacity int    `json:"capacity"`
	Status   string `json:"status"`
}
