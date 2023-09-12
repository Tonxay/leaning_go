package main

import (
	"fmt"
	"gomodApp/database"
	initialzers "gomodApp/initializers"
	"gomodApp/logs"
	"gomodApp/router"
	"log"

	"github.com/gofiber/fiber/v2"
)


func main(){
	app := fiber.New()
    initialzers.LoadEnvVariable()
    database.InitDB()
	logs.SetLoggerMiddleware(app)
	router.Router(app)
    fmt.Println("Successfully connected!")
	log.Fatal(app.Listen(":8081"))
}