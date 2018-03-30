package checkpoint

// PostSource is checkpoint for POST /source
type PostSource struct {
	URL      string `json:"url" binding:"required" validate:"url"`
	Category int64  `json:"category_id" binding:"required" validate:"min=1"`
}

// DelSource is checkpoint for DEL /source/:id
type DelSource struct {
	ID int64 `json:"id" validate:"min=1"`
}

// PutSource is checkpoint for PUT /source/:id
type PutSource struct {
	Alias    string `json:"alias"`
	Category int64  `json:"category_id"`
}

// GetSourceList is checkpoint for GET /source?category=xx
type GetSourceList struct {
	Category string `form:"category"`
}
