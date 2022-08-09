package database

import (
	"crosscheck-golang/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgres(c *config.Config) *sql.DB {
	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.DbConfig.Host,
		c.DbConfig.Port,
		c.DbConfig.User,
		c.DbConfig.Database,
		c.DbConfig.Password,
	)

	log.Println(connection)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println("Error Connect!")
		log.Fatal(err)
	}

	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Println("Error Ping!")
		log.Fatal(err)
	}

	log.Println("Connecting Database is successfully!")

	return db
}
