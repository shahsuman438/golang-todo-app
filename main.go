package main

import (
	// "log"
	"net/http"
	"todoapp/configs"
	"todoapp/controllers"
	// "todoapp/routes"
	// "github.com/gofiber/fiber/v2"
	// "github.com/gofiber/template/html"
)

// func main() {
// 	app := fiber.New(fiber.Config{
// 		Views: html.New("./templates", ".tpl"),
// 	})
// 	configs.ConnectDB()
// 	routes.TodoRoute(app)
// 	log.Fatal(app.Listen("localhost:3000"))
// }

// func main() {
// 	configs.ConnectDB()
// 	http.HandleFunc("/", controllers.Home)
// 	http.HandleFunc("/getall", controllers.GetATodo)
// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	configs.ConnectDB()
	http.HandleFunc("/todo", controllers.CreateTodo)
	http.HandleFunc("/gettodo", controllers.GetATodo)
	http.HandleFunc("/updatetodo", controllers.EditATodo)
	http.HandleFunc("/deletetodo", controllers.DeleteATodo)
	http.HandleFunc("/getalltodos", controllers.GetAllTodos)
	http.ListenAndServe("localhost:8080", nil)
}
