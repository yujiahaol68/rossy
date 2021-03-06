package post

import (
	"github.com/yujiahaol68/rossy/app/database"
	"github.com/yujiahaol68/rossy/app/entity"
)

func Insert(p *entity.Post) error {
	db := database.Conn()
	_, err := db.InsertOne(p)
	return err
}

func List(limit, offset int) ([]*entity.Post, error) {
	db := database.Conn()

	rl := make([]*entity.Post, 0)
	err := db.Cols("id", "desc", "source_id", "title", "link", "author", "create_at").Desc("create_at").Limit(limit, offset).Find(&rl)
	if err != nil {
		return nil, err
	}

	return rl, nil
}

func UnreadList(limit, offset int) ([]*entity.Post, error) {
	db := database.Conn()

	unread := make([]*entity.Post, 0)
	err := db.Cols("id", "desc", "source_id", "title", "link", "author", "create_at").Where("post.unread = 1").Desc("create_at").Limit(limit, offset).Find(&unread)
	if err != nil {
		return nil, err
	}

	return unread, nil
}

func SourceList(sourceID int64, limit, offset int) ([]*entity.Post, error) {
	db := database.Conn()

	srl := make([]*entity.Post, 0)
	err := db.Cols("id", "desc", "source_id", "title", "link", "author", "create_at").Where("post.source_id = ?", sourceID).Desc("create_at").Limit(limit, offset).Find(&srl)
	if err != nil {
		return nil, err
	}

	return srl, err
}

func DelBySource(source int64) error {
	db := database.Conn()
	_, err := db.Where("post.source_id = ?", source).Delete(&entity.Post{})
	return err
}

func MarkRead(id int64) error {
	_, err := database.Conn().Table(new(entity.Post)).ID(id).Update(map[string]interface{}{"unread": false})
	return err
}
