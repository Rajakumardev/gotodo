package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/rajakumardev/gotodo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb(config *Config) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", config.DBHOST, config.DBUSER, config.DBPASSWORD, config.PORT)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("[!]Failed to connect to database ...", err.Error())
		os.Exit(1)
	}

	log.Println("[*] Connected to database...")
	log.Println("[*] Running migration...")

	DB.AutoMigrate(&models.Todo{})

	log.Println("[*] intialized to database connection...")

}
