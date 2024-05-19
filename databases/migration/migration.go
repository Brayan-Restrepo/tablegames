package main

import (
	"gorm.io/gorm"
	"tablegames/config"
	"tablegames/databases"
	"tablegames/entities"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	tx := db.Connect().Begin()

	gamesMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func gamesMigration(tx *gorm.DB) {
	err := tx.Migrator().CreateTable(&entities.Game{})
	if err != nil {
		return
	}
}
