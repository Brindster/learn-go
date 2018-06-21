package model

import (
	"github.com/jinzhu/gorm"
	"os"

	"chrisbrindley.co.uk/service"

	// Import mysql dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	conn = os.Getenv("DB_CONN")
)

// NewDbConnection creates a new database connection
func NewDbConnection(c service.Container) (interface{}, error) {
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return db, nil
}
