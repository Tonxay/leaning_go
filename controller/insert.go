package controller

import (
	"gomodApp/database"
	"gomodApp/model"

	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)
func InsertData()fiber.Handler{
         insertStatement := 
		 `INSERT INTO sites (id, name)
          VALUES ($1, $2)`
	return func(c *fiber.Ctx) error {
		var  data model.Domain
		if err := c.BodyParser(&data); err != nil {
			return c.Status(fiber.StatusBadGateway).SendString(err.Error())
		}
	    database.DB.QueryRow(insertStatement,&data.Id,&data.Domain).Scan(&data.Id,&data.Domain)
		return c.JSON(data)
	}
}