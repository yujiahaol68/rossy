package checkpoint

// PostSource is checkpoint for POST /source
type PostSource struct {
	URL      string `json:"url" binding:"required" validate:"url"`
	Category int64  `json:"category_id" binding:"required" validate:"min=1"`
}

// DelSource is checkpoint for DEL /source/:id
type DelSource struct {
	ID int64 `form:"id" validate:"min=1"`
}

// PutSource is checkpoint for PUT /source:id
type PutSource struct {
}

// GetSourceListBy is checkpoint for GET /source?category=xx
type GetSourceListBy struct {
	Category int64 `form:"category" binding:"required"`
}
