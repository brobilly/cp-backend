package controller

// func SendMail(c *fiber.Ctx) error {
// 	var emailData models.EmailData
// 	if err := c.BodyParser(&emailData); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
// 	}

// 	m := mail.NewMessage()
// 	m.SetHeader("From", "daudmanuwu@gmail.com")
// 	m.SetHeader("To", emailData.To)
// 	m.SetHeader("Subject", emailData.Subject)

// 	templatePath := filepath.Join("templates", "email_template.html")
// 	tmpl, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		log.Println(err)
// 		return c.Status(500).JSON(fiber.Map{"error": "Gagal memuat template email"})
// 	}

// 	var bodyContent strings.Builder
// 	if err := tmpl.Execute(&bodyContent, emailData); err != nil {
// 		log.Println(err)
// 		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengeksekusi template"})
// 	}

// 	m.SetBody("text/html", bodyContent.String())

// 	qrImageBase64 := emailData.QRImage
// 	qrImageData, err := base64.StdEncoding.DecodeString(qrImageBase64)
// 	if err != nil {
// 		log.Println(err)
// 		return c.Status(500).JSON(fiber.Map{"error": "Gagal mendecode gambar QR code"})
// 	}

// 	qrImagePath := filepath.Join("temp", "qr_code.jpg")
// 	err = os.WriteFile(qrImagePath, qrImageData, 0644)
// 	if err != nil {
// 		log.Println(err)
// 		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan gambar QR code"})
// 	}

// 	m.Attach(qrImagePath)

// 	d := mail.NewDialer("smtp.gmail.com", 587, "daudmanuwu@gmail.com", "mqfznxvehmitlxir")

// 	if err := d.DialAndSend(m); err != nil {
// 		log.Println(err)
// 		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengirim email"})
// 	}

// 	if err := os.Remove(qrImagePath); err != nil {
// 		log.Println(err)
// 	}

// 	fmt.Println("Email berhasil dikirim.")
// 	return c.JSON(fiber.Map{"message": "Email berhasil dikirim"})
// }
