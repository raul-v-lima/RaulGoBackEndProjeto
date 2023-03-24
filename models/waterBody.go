package models

import "gorm.io/gorm"

type WaterBody struct {
	gorm.Model
	ID   uint64  `json:"id" gorm: "auto_increment;primary_key"`
	Name *string `json:"name" gorm: "text;not null;default:null`

	EnvironmentID uint
}
