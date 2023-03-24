package models

import "gorm.io/gorm"

type Environment struct {
	gorm.Model
	ID   uint    `json:"id" gorm: "auto_increment;primary_key"`
	Name *string `json:"name" gorm: "text;not null;default:null`

	Assets      []Asset
	WaterBodies []WaterBody
	Terrains    []Terrain
	//Creatures          []Creature
	Decoys             []Decoy
	EndemicCreatures   []EndemicCreature
	EndemicVegetations []EndemicVegetation
	Landscapes         []Landscape
	Vegetations        []Vegetation
}
