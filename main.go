package main

import (
	"log"
	"task02/config"
	"task02/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewRouter() *fiber.App {
	router := fiber.New()
	router.Route("api/:user_id", func(router fiber.Router) {
		router.Delete("", func(c *fiber.Ctx) error {
			user_id, err := c.ParamsInt("user_id")
			new_err := model.Error{}
			if err != nil {
				new_err.Message = "Can't parse the ID from the query"
				c.JSON(&new_err)
				return nil
			}
			var person model.Person
			ptr := DB.Find(&person, user_id)
			if ptr == nil {
				new_err.Message = "No row with found with that id"
				c.JSON(&new_err)
				return nil
			}
			person_old := person
			DB.Unscoped().Delete(&person, user_id)
			c.JSON(person_old)
			return nil
		})

		router.Put("", func(c *fiber.Ctx) error {
			user_id, err := c.ParamsInt("user_id")
			new_err := model.Error{}
			if err != nil {
				new_err.Message = "Can't parse the ID from the query"
				c.JSON(&new_err)
				return nil
			}
			new_person := model.Person{}
			err = c.BodyParser(&new_person)
			if err != nil {
				new_err.Message = "Failed to parse body"
				c.JSON(&new_err)
				return nil
			}
			person := model.Person{}
			ptr := DB.Find(&person, user_id)
			if ptr == nil {
				new_err.Message = "No row with found with that id"
				c.JSON(&new_err)
				return nil
			}
			old_person := person
			person.Name = new_person.Name
			DB.Save(&person)
			c.JSON(old_person)
			return nil
		})

		router.Patch("", func(c *fiber.Ctx) error {
			user_id, err := c.ParamsInt("user_id")
			new_err := model.Error{}
			if err != nil {
				new_err.Message = "Can't parse the ID from the query"
				c.JSON(&new_err)
				return nil
			}
			new_person := model.Person{}
			err = c.BodyParser(&new_person)
			if err != nil {
				new_err.Message = "Failed to parse body"
				c.JSON(&new_err)
				return nil
			}
			person := model.Person{}
			ptr := DB.Find(&person, user_id)
			if ptr == nil {
				new_err.Message = "No row with found with that id"
				c.JSON(&new_err)
				return nil
			}
			old_person := person
			person.Name = new_person.Name
			DB.Save(&person)
			c.JSON(old_person)
			return nil
		})

		router.Get("", func(c *fiber.Ctx) error {
			user_id, err := c.ParamsInt("user_id")
			new_err := model.Error{}
			if err != nil {
				new_err.Message = "Can't parse the ID from the query"
				c.JSON(&new_err)
				return nil
			}
			var person = model.Person{}
			ptr := DB.Find(&person, user_id)
			if ptr == nil {
				new_err.Message = "No row with found with that id"
				c.JSON(&new_err)
				return nil
			}
			c.JSON(person)
			return nil
		})

	})

	router.Route("/api", func(router fiber.Router) {
		router.Get("", func(c *fiber.Ctx) error {
			var people = []model.Person{}
			DB.Find(&people)
			c.JSON(people)
			return nil
		})

		router.Post("", func(c *fiber.Ctx) error {
			person := model.Person{}
			err := c.BodyParser(&person)
			new_err := model.Error{}
			if err != nil {
				new_err.Message = "Failed to parse body"
				c.JSON(&new_err)
				return nil
			}
			DB.Create(&person)
			c.JSON(person)
			return nil
		})

	})

	router.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "message",
			"message": "Alive",
		})
	})

	return router
}

func main() {
	app := fiber.New()
	DB = config.ConnectionDb(&model.Person{})
	app.Mount("/", NewRouter())
	log.Fatal(app.Listen(":8080"))
}
