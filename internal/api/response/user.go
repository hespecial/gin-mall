package response

type UserInfoUpdateResp struct{}

type UserPasswordChangeResp struct{}

type ShowUserInfoResp struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Status   string `json:"status"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	// Money    string `json:"money"`
}

type UploadAvatarResp struct{}
