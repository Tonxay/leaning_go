package controller

import (
	"gomodApp/database"
	"gomodApp/model"

	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)



func Guerydata ()fiber.Handler{
	var datas []model.Domain
	results, err :=  database.DB.Query("SELECT id,name FROM sites")
	for results.Next() {
      var domain model.Domain
		// for each row, scan the result into our tag composite object
		err = results.Scan(&domain.Id, &domain.Domain)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute	
     datas = append(datas,domain)
	}
    defer results.Close()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	  
   return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(datas)
	}
}
