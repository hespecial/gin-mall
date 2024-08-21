package e

type Code int

const (
	Success       Code = iota // 响应成功
	InvalidParams             // 参数错误
	UnknownError              // 未知错误

	ErrorUserExists        // 用户已存在
	ErrorEncryptPassword   // 密码加密错误
	ErrorEncryptMoney      // 金额加密错误
	ErrorCreateUser        // 创建用户错误
	ErrorAccountInvalid    // 用户名或密码错误
	ErrorGetUserByID       // 根据id获取用户失败
	ErrorUpdateUser        // 更新用户失败
	ErrorIncorrectPassword // 密码错误
	ErrorUploadAvatar      // 头像上传错误

	ErrorGenerateToken // token生成错误
	ErrorParseToken    // token解析错误
	ErrorContextValue  // 上下文值传递错误

	ErrorUploadFile     // 文件上传错误
	ErrorFileError      // 文件错误
	ErrorOSSUploadError // OSS文件上传错误
	ErrorFileType       // 文件类型错误

	ErrorSendEmail            // 发送邮件错误
	ErrorUpdateEmail          // 更新邮箱错误
	ErrorEmailLinkExpire      // 邮件确认链接已过期
	ErrorSendEmailTooFrequent // 邮件发送操作频繁
)
