package database

import (
	"fmt"
	"log"
	"os"

	"projetoRaul/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Environment{}, &models.Character{}, &models.Player{},
		&models.Item{}, &models.Asset{}, &models.Building{}, &models.Climate{},
		&models.Creature{}, &models.Decoy{}, &models.EndemicCreature{}, &models.EndemicVegetation{},
		&models.Landscape{}, &models.Terrain{}, &models.WaterBody{}, &models.Vegetation{})

	DB = Dbinstance{
		Db: db,
	}
}
