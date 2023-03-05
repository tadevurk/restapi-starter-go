package handlers

import (
	"github.com/divrhino/divrhino-trivia/database"
	"github.com/divrhino/divrhino-trivia/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Rest API v2.0 with Golang and Fiber")
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func UpdateFact(c *fiber.Ctx) error {
	id := c.Params("id")
	fact := new(models.Fact)
	database.DB.Db.Where("id = ?", id).First(&fact)

	if fact.Question == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "Fact not found",
		})
	}

	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Save(&fact)

	return c.Status(200).JSON(fact)
}

func DeleteFact(c *fiber.Ctx) error {
	id := c.Params("id")
	fact := new(models.Fact)
	database.DB.Db.Where("id = ?", id).Delete(&fact)

	return c.Status(200).JSON(fiber.Map{
		"message": "Fact deleted",
	})
}

func GetFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)
	return c.Status(200).JSON(facts)
}
