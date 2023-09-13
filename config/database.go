package config

import (
	"log"
	"task02/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectionDb(persondto *model.Person) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running AutoMigration")
	db.AutoMigrate(persondto)
	return db
}

func OpenConnectionDb(persondto *model.Person, db_filename string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(db_filename), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running AutoMigration")
	db.AutoMigrate(persondto)
	return db
}
