package database

import (
	"fmt"
	"github.com/ArtiomStartev/todo-list-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "artiomstartev"
	password = "postgres"
	dbName   = "todo-list-api"
)

var (
	DB  *gorm.DB
	dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
)

func DBConn() error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return err
	}
	DB = db

	if err = db.AutoMigrate(&models.Task{}); err != nil {
		fmt.Println("Error migrating database: ", err)
		return err
	}

	return nil
}
