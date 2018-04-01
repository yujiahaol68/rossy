package endpoint

type UnreadSourceList map[string][]Source

type Source struct {
	ID          int64  `json:"source_id"`
	Category    int64  `json:"category_id"`
	Alias       string `json:"alias"`
	UnreadCount int64  `json:"unread_count"`
}
