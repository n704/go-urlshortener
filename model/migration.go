package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// InitialMigration migrate all models
func InitialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&ShortLink{})
}
