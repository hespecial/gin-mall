package response

type Category struct {
	ID   uint
	Name string
}

type GetCategoryListResp struct {
	List  []*Category
	Total int
}
