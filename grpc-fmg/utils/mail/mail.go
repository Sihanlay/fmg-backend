package mailUtils

import (
	"crypto/tls"
	"grpc-demo/utils"
	"github.com/go-gomail/gomail"
)

func Send(token string, target ...string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", utils.GlobalConfig.Server.Mail.Username)
	m.SetHeader("To", target...)

	m.SetHeader("Subject", "捣蒜官方邮件")
	m.SetBody("text/html", generateBody(token))
	d := gomail.NewDialer(
		utils.GlobalConfig.Server.Mail.SmtpHost, 465,
		utils.GlobalConfig.Server.Mail.Username, utils.GlobalConfig.Server.Mail.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
