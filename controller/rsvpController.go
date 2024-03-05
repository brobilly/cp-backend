package controller

import (
	"campus-api/database"
	"campus-api/models"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mail.v2"
)

func CreateReservation(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to parse JSON payload",
		})
	}

	var emailData models.EmailData
	if err := c.BodyParser(&emailData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to parse email JSON payload",
		})
	}

	if file, err := c.FormFile("qrimage"); err == nil {
		src, err := file.Open()
		if err != nil {
			c.Status(500)
			return c.JSON(fiber.Map{
				"message": "Unable to open the uploaded file",
			})
		}
		defer src.Close()

		imageBytes, err := io.ReadAll(src)
		if err != nil {
			c.Status(500)
			return c.JSON(fiber.Map{
				"message": "Unable to read the uploaded file",
			})
		}

		user.QRImage = imageBytes
	}

	emailData.Name = user.Name
	emailData.QRImage = user.QRImage
	emailData.To = user.Email

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create reservation",
		})
	}

	if err := database.DB.Create(&emailData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert data into email_data table",
		})
	}

	if err := sendReservationEmail(emailData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to send email",
		})
	}

	return c.JSON(fiber.Map{
		"message":     "Reservation created successfully and email sent",
		"reservation": user,
	})
}

func sendReservationEmail(emailData models.EmailData) error {
	m := mail.NewMessage()
	m.SetHeader("From", "daudmanuwu@gmail.com")
	m.SetHeader("To", emailData.To)
	m.SetHeader("Subject", "Campus Party Reservation")

	templatePath := filepath.Join("templates", "email_template.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println(err)
		return err
	}

	var bodyContent strings.Builder
	if err := tmpl.Execute(&bodyContent, emailData); err != nil {
		log.Println(err)
		return err
	}

	m.SetBody("text/html", bodyContent.String())

	qrImageData := emailData.QRImage
	qrImagePath := filepath.Join("temp", "qr_code.jpg")
	err = os.WriteFile(qrImagePath, qrImageData, 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	m.Attach(qrImagePath)

	d := mail.NewDialer("smtp.gmail.com", 587, "daudmanuwu@gmail.com", "mqfznxvehmitlxir")

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return err
	}

	if err := os.Remove(qrImagePath); err != nil {
		log.Println(err)
	}

	fmt.Println("Email berhasil dikirim.")
	return nil
}
