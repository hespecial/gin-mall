package email

import "testing"

func TestSender_SendEmail(t *testing.T) {
	LoadEmailConfig(&Config{
		Host:     "smtp.qq.com",
		Port:     25,
		Alias:    "gin-mall",
		Username: "1478488313@qq.com",
		Password: "qmixpqclxmorbaei",
	})
	sender := NewSender()
	err := sender.SendEmail("1478488313@qq.com", "hello", "hello world")
	if err != nil {
		t.Fatal(err)
	}
}
