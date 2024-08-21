package initialize

import (
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/pkg/jwt"
	"time"
)

func InitJWT() {
	const day = 24 * time.Hour
	jwt.LoadJWTConfig(&jwt.Config{
		Secret:          []byte(global.Config.Jwt.Secret),
		Issuer:          global.Config.Jwt.Issuer,
		AccessTokenTTL:  time.Duration(global.Config.Jwt.AccessTokenTTl) * day,
		RefreshTokenTTL: time.Duration(global.Config.Jwt.RefreshTokenTTl) * day,
	})
}
