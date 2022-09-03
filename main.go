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
		Views: html.New("./templates", ".tpl"),
	})
	configs.ConnectDB()
	routes.TodoRoute(app)
	log.Fatal(app.Listen("localhost:3000"))
}

// func main() {
// 	configs.ConnectDB()
// 	http.HandleFunc("/", controllers.Home)
// 	http.HandleFunc("/getall", controllers.GetATodo)
// 	http.ListenAndServe(":8080", nil)
// }
