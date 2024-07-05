package mail

import (
	"gateway/models"
	u "gateway/utils"
)

func SendEmailToUser(data interface{}) {
	user := data.(*models.Users)

	file := "./templates/emails/welcome.html"
	if content, err := u.ParseHTML(file, user); err == nil {
		u.Mailer("Welcome!", content, user.Email)
	}
}
