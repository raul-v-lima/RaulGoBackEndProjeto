package models

import "gorm.io/gorm"

type Decoy struct {
	gorm.Model
	ID            uint    `json:"id" gorm: "auto_increment;primary_key"`
	Name          *string `json:"name" gorm: "text;not null;default:null`
	EnvironmentID uint
}
