package response

type Favorite struct {
	ID       uint    `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
}

type GetFavoriteListResp struct {
	List []*Favorite `json:"list"`
}

type AddFavoriteResp struct{}

type DeleteFavoriteResp struct{}
