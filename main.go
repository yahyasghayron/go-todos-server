package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

type Todo struct {
	ID   int    `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"name": "yahya",
		})
	})
	// create todo
	app.Post("api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		return c.JSON(todos)

	})
	// marck todo as done 

	app.Patch("api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid Id")
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = true
				break
			}
		}

		return c.JSON(todos)
	})
	
	// marck todo as undone
		app.Patch("api/todos/:id/undone", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid Id")
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = false
				break
			}
		}

		return c.JSON(todos)
	})
	// get all todos
	app.Get("api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})
	// get  todo
	app.Get("api/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid Id")
		}

		for i, t := range todos{
			if t.ID == id  {
				return c.JSON(todos[i])
			}
		}
		return c.Status(404).SendString("todo not found")
	})
	// delete todo
	app.Delete("api/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid Id")
		}

		for i, t := range todos{
			if t.ID == id  {
				todos = append(todos[:i], todos[i+1:]...)
			}
		}
		return c.JSON(todos)
	})
	// update todo
	app.Patch("api/todos/:id", func(c *fiber.Ctx) error {
		id , err :=c.ParamsInt("id")

		todo := &Todo{}

		if errb := c.BodyParser(todo); errb != nil {
			return errb
		}
		if err != nil {
			return c.Status(401).SendString("Invalid Id")
		}

		for i , t := range todos{
			if t.ID == id {
				todos[i].Todo = todo.Todo
				return c.JSON(todos)
			}
		}

		return c.Status(404).SendString("todo Not Found")
	})


	log.Fatal(app.Listen(":4050"))
}
