package checkpoint

type PageArg struct {
	From int `form:"offset" binding:"exists" validate:"gte=0"`
	Size int `form:"limit" binding:"required" validate:"gte=0"`
}
