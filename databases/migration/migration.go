package main

import (
	"log"
	"thanapatjitmung/go-test-komgrip/config"
	"thanapatjitmung/go-test-komgrip/databases"
	"thanapatjitmung/go-test-komgrip/entities"

	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	mariaDb := databases.NewMariaDatabase(conf.MariaDB)
	
	tx := mariaDb.Begin()

	beerMigration(mariaDb)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(tx.Error)
	}
	log.Println("Table created successfully!")
}

func beerMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Beer{})
}
