package constant

import "time"

const (
	UploadModeLocal = "local"
	UploadModeOSS   = "oss"

	UserID   = "userID"
	Username = "username"

	AvatarTypeJPEG = "image/jpeg"
	AvatarTypePNG  = "image/png"

	ValidEmailAddress = "api/v1/user/email/valid"
	EmailSubject      = "gin-mall"
	BindEmailBody     = "您现在正在绑定邮箱，如确认是您本人操作，请点击链接：<a>%s<a>"
	EmailTokenKey     = "token"
	EmailLimiterR     = time.Hour
	EmailLimiterB     = 3

	CartKey       = "cart"
	CartKeyExpire = time.Hour * 24
)
