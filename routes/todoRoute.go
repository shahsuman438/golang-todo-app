package routes

import (
	"todoapp/controllers"

	"github.com/gofiber/fiber/v2"
)

func TodoRoute(app *fiber.App) {
	//All routes related to todos comes here
	app.Post("/todo", controllers.CreateTodo)
	app.Get("/todo/:todoId", controllers.GetATodo)
	app.Put("/todo/:todoId", controllers.EditATodo)
	app.Delete("/todo/:todoId", controllers.DeleteATodo)
	app.Get("/todos", controllers.GetAllTodos)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{
		   "Title": "Hello, World!",
		})
	  })
}
