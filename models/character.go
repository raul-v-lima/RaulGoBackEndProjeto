package models

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	ID        uint    `json:"id" gorm: "auto_increment;primary_key"`
	Name      *string `json:"name" gorm: "size:255";not null;default:null`
	Vigor     int     `json:"vigor" gorm: "not null;default:null`
	Dexterity int     `json:"dexterity" gorm: "not null;default:null`
	Xp        int     `json:"xp" gorm: "not null;default:null`
	Level     int     `json:"level" gorm: "not null;default:null`
	Mana      int     `json:"mana" gorm: "not null;default:null`

	Players []Player
	Items   []Item
}
