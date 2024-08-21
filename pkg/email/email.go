package email

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
)

type Config struct {
	Host     string
	Port     int
	Alias    string
	Username string
	Password string
}

var conf *Config

func LoadEmailConfig(c *Config) {
	conf = c
}

type Sender struct {
}

func NewSender() *Sender {
	return &Sender{}
}

func (*Sender) SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", conf.Alias, conf.Username))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(conf.Host, conf.Port, conf.Username, conf.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
