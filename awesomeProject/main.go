package main

import (
	"awesomeProject/models"
	_ "awesomeProject/models"
	_ "strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
)

type User struct {
	Id         string
	Firstname  string
	SecondName string
}

var listaUsuarios []User

func handleUser(c *fiber.Ctx) error {
	user := User{
		uuid.NewString(),
		"Manuel",
		"lopez",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.Id = uuid.NewString()
	listaUsuarios = append(listaUsuarios, user)
	return c.Status(fiber.StatusOK).JSON(user)
}

func handleGetAllUsers(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(listaUsuarios)
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5174/",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())

	app.Get("/", func(f *fiber.Ctx) error {
		return f.Status(fiber.StatusOK).JSON(models.EuroData())
	})

	app.Get("/prueba", func(c *fiber.Ctx) error {
		return c.SendString("Probando")
	})

	app.Get("/user", handleUser)

	userGroup := app.Group("/users")

	userGroup.Get("", handleGetAllUsers)
	userGroup.Post("", handleCreateUser)

	app.Listen(":3000")
}
