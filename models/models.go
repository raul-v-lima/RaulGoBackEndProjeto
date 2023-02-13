package models

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	ID         uint64  `json:"id" gorm: "auto_increment;primary_key"`
	Name       *string `json:"name" gorm: "text;not null;default:null`
	Vigor      *int32  `json:"vigor"gorm: "text;not null;default:null`
	Empiricism *int32  `json:"empiricism" gorm:"text;not null;default:null`
	Dexterity  *int32  `json:"dexterity" gorm:"text;not null;default:null`
	Mana       *int32  `json:"mana" gorm:"text;not null;default:null`
	Xp         *int32  `json:"xp" gorm:"text;not null;default:null`
	Level      *int32  `json:"level" gorm:"text;not null;default:null`
}
