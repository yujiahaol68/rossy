package entity

import "time"

// Source entity
type Source struct {
	ID           int64     `xorm:"pk autoincr 'id'" json:"id"`
	Category     int64     `xorm:"notnull 'category_id'" json:"category_id"`
	URL          string    `xorm:"notnull 'url'" json:"url"`
	ETag         string    `xorm:"'etag'" json:"etag"`
	LastModified string    `json:"last_modified"`
	Alias        string    `xorm:"notnull" json:"alias"`
	Kind         string    `xorm:"notnull"`
	CreateAt     time.Time `xorm:"created" json:"create_at"  time_format:"2006-01-02 15:04:05"`
	DeleteAt     time.Time `xorm:"deleted"`
	Updated      time.Time `xorm:"updated"`
}

// Post entity
type Post struct {
	ID       int64     `xorm:"pk autoincr 'id'" json:"id"`
	Desc     string    `xorm:"notnull" json:"summary"`
	From     int64     `xorm:"notnull 'source_id'" json:"source_id"`
	Title    string    `xorm:"notnull" json:"title"`
	Link     string    `xorm:"notnull unique" json:"link"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	CreateAt time.Time `json:"create_at"  time_format:"2006-01-02 15:04:05"`
	DeleteAt time.Time `xorm:"deleted"`
	Unread   bool      `json:"unread"`
}

// Category entity
type Category struct {
	ID      int64     `xorm:"pk autoincr 'id'" json:"id"`
	Name    string    `xorm:"notnull unique" json:"name"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
