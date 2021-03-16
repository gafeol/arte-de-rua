package classes

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

type Art struct {
	gorm.Model
	Frase  string `json:"frase"`
	ImgURL string `json:"imgURL"`
}

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
		if db == nil {
			println("FODEU DB JA EH NIL")
		}
	})
	return db
}

func (art Art) Create() error {
	db := initDB()
	return db.Create(art).Error
}

func FindArt(id string) (Art, error) {
	db := initDB()
	var art Art
	println("Vamo que vamo busca")
	err := (*db).First(&art, id).Error
	println("saiu da busca aqi")
	println("Finding art with id ", id, "got actually ", err)
	return art, err
}
