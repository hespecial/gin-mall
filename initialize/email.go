package initialize

import (
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/pkg/email"
)

func InitEmail() {
	email.LoadEmailConfig(&email.Config{
		Host:     global.Config.Email.Host,
		Port:     global.Config.Email.Port,
		Alias:    global.Config.Email.Alias,
		Username: global.Config.Email.Username,
		Password: global.Config.Email.Password,
	})
}
