package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "200320052010meir"
	dbname   = "qadam"
)

var Database *gorm.DB

func ConnectDB() {
	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	Database, err = gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
	fmt.Println("We connected successfully!")
}
