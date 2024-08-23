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

type BindEmailResp struct{}

type UnbindEmailResp struct{}

type ValidEmailResp struct{}

type UserFollowResp struct{}

type UserUnfollowResp struct{}

type Follow struct {
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

type UserFollowingListResp struct {
	Following []*Follow `json:"following"`
}

type UserFollowerListResp struct {
	Follower []*Follow `json:"follower"`
}
