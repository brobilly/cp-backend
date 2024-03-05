package controller

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"campus-api/database"
	"campus-api/models"
	"campus-api/util"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9. %+\-]+@[a-z0-9. %+\-]+\.[a-z0-9. %+\-]`)
	return Re.MatchString(email)
}

func RegisterAdmin(c *fiber.Ctx) error {
	var data map[string]interface{}
	var adminData models.Admin
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 6 character",
		})

	}
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Email Address",
		})
	}
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&adminData)
	if adminData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email Already exist",
		})
	}
	admin := models.Admin{
		Name:  data["name"].(string),
		Phone: data["phone"].(string),
		Email: strings.TrimSpace(data["email"].(string)),
	}
	admin.SetPassword(data["password"].(string))
	err := database.DB.Create((&admin))
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"admin":   admin,
		"message": "Account created successfully",
	})
}

func LoginAdmin(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	var admin models.Admin
	database.DB.Where("email=?", data["email"]).First(&admin)
	if admin.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Email Address doesn't exist, create an account",
		})
	}
	if err := admin.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}
	token, err := util.GenerateJwt(strconv.Itoa(int(admin.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "you have succesfully login",
		"admin":   admin,
		"cookie":  cookie,
	})
}

func UserRegister(c *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 6 character",
		})

	}
	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Email Address",
		})
	}
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email Already exist",
		})
	}
	user := models.User{
		Name:  data["name"].(string),
		Phone: data["phone"].(string),
		Email: strings.TrimSpace(data["email"].(string)),
	}
	user.SetPassword(data["password"].(string))
	err := database.DB.Create((&user))
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Account created successfully",
	})
}

func UserLogin(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Email Address doesn't exist, create an account",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "you have succesfully login",
		"user":    user,
		"cookie":  cookie,
	})
}

type Claims struct {
	jwt.StandardClaims
}
