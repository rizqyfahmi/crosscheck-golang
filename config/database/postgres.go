package database

import (
	"crosscheck-golang/config"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgres(c *config.Config) *sqlx.DB {
	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.DbConfig.Host,
		c.DbConfig.Port,
		c.DbConfig.User,
		c.DbConfig.Database,
		c.DbConfig.Password,
	)

	db, err := sqlx.Open("postgres", connection)
	if err != nil {
		log.Println("Error connect!")
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Println("Error ping!")
		log.Fatal(err)
	}

	log.Println("Database successfully connected!")
	return db
}
