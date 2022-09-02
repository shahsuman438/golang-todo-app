package main

import (
	"log"
	"todoapp/configs"
	"todoapp/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	app := fiber.New(fiber.Config{
        Views: html.New("./templates", ".tmpl"),
    })
	configs.ConnectDB()
	routes.TodoRoute(app)
	log.Fatal(app.Listen(":3000"))
}
