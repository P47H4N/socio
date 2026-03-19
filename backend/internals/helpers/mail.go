package helpers

import (
	"fmt"
	"log"
	"net/smtp"
)



func SendEmail(toEmail, token, mailtype string) {
	from := GMAIL
	password := APP_PASSWORD
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	url := fmt.Sprintf("%s/confirm?email=%s&token=%s", BASE_URL, toEmail, token)
	var subject, body string
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	
	if mailtype=="password_reset" {
		subject = "Subject: Password Reset Request - Socio\r\n"
		body = resetPasswordBody(url)
	}
	
	message := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, message)
	if err != nil {
		log.Fatalln(err)
	}
}

func resetPasswordBody(url string) string {
	return fmt.Sprintf(`
    <html>
        <body style="font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #f4f4f4; padding: 20px;">
            <div style="max-width: 600px; margin: 0 auto; background-color: #ffffff; border-radius: 10px; overflow: hidden; box-shadow: 0 4px 10px rgba(0,0,0,0.1);">
                <div style="background-color: #007bff; padding: 20px; text-align: center; color: white;">
                    <h1 style="margin: 0;">Socio</h1>
                </div>
                <div style="padding: 30px; line-height: 1.6; color: #333;">
                    <h2 style="color: #007bff; text-align: center;">Password Reset Request</h2>
                    <p>Hello,</p>
                    <p>We received a request to reset the password for your Socio account. Click the button below to proceed:</p>
                    
                    <div style="text-align: center; margin: 30px 0;">
                        <a href="%s" style="display: inline-block; padding: 12px 30px; background-color: #007bff; color: white; text-decoration: none; border-radius: 5px; font-weight: bold; font-size: 16px;">Reset Password</a>
                    </div>

                    <p style="font-size: 14px; color: #666;">If the button above does not work, please copy and paste the following link into your web browser:</p>
                    <p style="word-break: break-all; font-size: 13px;"><a href="%s" style="color: #007bff;">%s</a></p>
                    
                    <hr style="border: none; border-top: 1px solid #eee; margin: 20px 0;">
                    <p style="font-size: 12px; color: #999; text-align: center;">
                        This link is valid for <b>15 minutes</b>. If you did not request this, you can safely ignore this email.
                    </p>
                </div>
                <div style="background-color: #f8f9fa; padding: 15px; text-align: center; font-size: 12px; color: #777;">
                    &copy; 2026 Socio. All rights reserved.
                </div>
            </div>
        </body>
    </html>`, url, url, url)
}