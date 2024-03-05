package controller

import (
	"campus-api/database"
	"campus-api/models"
	"campus-api/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	var user models.User

	if err := database.DB.Model(&user).Where("id = ?", id).Find(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve user from the database",
		})
	}

	return c.JSON(fiber.Map{
		"Id":   user.Id,
		"User": user,
	})
}

func GetUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	Id, err := util.ParseJwt(cookie)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	user := models.User{
		Id: uint(Id),
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to parse request body",
		})
	}

	var users []models.User
	database.DB.Model(&users).Where("id = ?", user.Id).Find(&users)

	return c.JSON(fiber.Map{
		"Id":   user.Id,
		"User": users,
	})
}
