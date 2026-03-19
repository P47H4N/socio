package helpers

import (
	"net/smtp"
	"fmt"
)

func SendEmail(toEmail string, token string) error {
	from := GMAIL
	password := APP_PASSWORD
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Password Reset - Socio\n"
	body := fmt.Sprintf("Your password reset token is: %s\n\nThis token is valid for 15 minutes.", token)
	message := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, message)
	if err != nil {
		return err
	}
	return nil
}