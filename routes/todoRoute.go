package routes

import (
	"net/http"
	"todoapp/controllers"

	// "github.com/gofiber/fiber/v2"
)

func TodoRoute(app *http.ServeMux) {
	//All routes related to todos comes here
	app.HandleFunc("/todo", controllers.CreateTodo)
	// app.Get("/todo/:todoId", controllers.GetATodo)
	// app.Put("/todo/:todoId", controllers.EditATodo)
	// app.Delete("/todo/:todoId", controllers.DeleteATodo)
	// app.Get("/todos", controllers.GetAllTodos)
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Render("Todo-home", fiber.Map{
	// 		"Title": "Golang Todo App",
	// 	})
	// })
}
