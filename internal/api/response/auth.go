package response

type AuthRegisterResp struct{}

type AuthLoginResp struct {
	Nickname     string `json:"nickname"`
	Avatar       string `json:"avatar"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
