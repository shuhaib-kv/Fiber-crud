package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuhaib-kv/Fiber-crud.git/db"
)

func main() {
	app := fiber.New()
	dbErr := db.InitDB()

	if dbErr != nil {
		panic(dbErr)
	}

	list := app.Group("/list")

	list.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Todo-List-API Tutorial :)")
	}) // "/" - Default route to return the given string.

}
