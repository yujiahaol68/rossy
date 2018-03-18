package entity

import "time"

// Source entity
type Source struct {
	ID           int64     `xorm:"pk autoincr 'id'" json:"id"`
	Category     int64     `xorm:"notnull 'category_id'" json:"category"`
	URL          string    `xorm:"notnull unique 'url'" json:"url"`
	ETag         string    `xorm:"'etag'" json:"etag"`
	LastModified string    `json:"last_modified"`
	Alias        string    `xorm:"notnull unique" json:"alias"`
	CreateAt     time.Time `json:"create_at"  time_format:"2006-01-02 15:04:05"`
	Updated      time.Time `xorm:"updated"`
}

// Post entity
type Post struct {
	ID       int64     `xorm:"pk autoincr 'id'" json:"id"`
	From     int64     `xorm:"notnull 'source_id'"`
	Title    string    `xorm:"notnull" json:"title"`
	Link     string    `xorm:"notnull" json:"link"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	CreateAt time.Time `json:"create_at"  time_format:"2006-01-02 15:04:05"`
	Unread   bool      `json:"unread"`
}

// Category entity
type Category struct {
	ID      int64     `xorm:"pk autoincr 'id'" json:"id"`
	Name    string    `xorm:"notnull unique" json:"name"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}