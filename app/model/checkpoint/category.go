package checkpoint

type PostCategory struct {
	Name string `json:"name" binding:"required"`
}
