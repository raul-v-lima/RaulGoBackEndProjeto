package models

import "gorm.io/gorm"

type Creature struct {
	gorm.Model
	ID          uint    `json:"id" gorm: "auto_increment;primary_key"`
	Name        *string `json:"name" gorm: "text;not null;default:null`
	CharacterID uint
}
