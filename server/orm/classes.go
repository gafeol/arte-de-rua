package classes

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once
var db *gorm.DB = nil

func initDB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: "host=localhost user=artederua password=SuperSecret dbname=artederua port=5432 sslmode=disable",
		}), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		if err = db.AutoMigrate(&Art{}); err != nil {
			panic(err)
		}
		if err = db.AutoMigrate(&Artist{}); err != nil {
			panic(err)
		}
	})
	return db
}

type Art struct {
	gorm.Model
	ID     uint64
	Frase  string
	ImgURL string
}

func (art *Art) Create() error {
	db := initDB()
	return db.Create(art).Error
}

func AllArts() ([]Art, error) {
	db := initDB()
	var arts []Art
	err := db.Find(&arts).Error
	return arts, err
}

func FindArt(id string) (Art, error) {
	db := initDB()
	var art Art
	err := db.First(&art, id).Error
	return art, err
}

type Artist struct {
	gorm.Model
	ID   uint64
	Nome string
}

func (artist *Artist) Create() error {
	db := initDB()
	return db.Create(artist).Error
}

func FindArtist(id string) (Artist, error) {
	db := initDB()
	var artist Artist
	err := db.First(&artist, id).Error
	return artist, err
}
