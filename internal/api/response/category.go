package response

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetCategoryListResp struct {
	List  []*Category `json:"list"`
	Total int         `json:"total"`
}
