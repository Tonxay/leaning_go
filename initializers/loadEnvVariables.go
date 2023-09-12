package initialzers

import (
	"log"

	"github.com/joho/godotenv"
)

// Load .env file found in this project
func LoadEnvVariable(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err.Error())
		log.Fatal("Error loading .env file")
	}
}