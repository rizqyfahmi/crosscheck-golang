package database

import (
	"crosscheck-golang/config"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewPostgres(c *config.Config) *gorm.DB {
	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.DbConfig.Host,
		c.DbConfig.Port,
		c.DbConfig.User,
		c.DbConfig.Database,
		c.DbConfig.Password,
	)

	db, err := gorm.Open("postgres", connection)
	if err != nil {
		log.Println("Error Connect!")
		log.Fatal(err)
	}

	defer db.Close()

	return db
}
