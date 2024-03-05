package models

type EmailData struct {
	ID      uint   `json:"id"`
	To      string `json:"to"`
	Name    string `json:"name"`
	QRImage []byte `json:"qrimage"`
}
