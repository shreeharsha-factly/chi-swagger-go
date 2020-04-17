package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB db setup
var DB *gorm.DB

// SetupDB is database setuo
func SetupDB() {
	fmt.Println("connecting to database ...")
	connStr := "user=postgres dbname=data_portal host=localhost sslmode=disable password=postgres"
	var err error
	DB, err = gorm.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

}
