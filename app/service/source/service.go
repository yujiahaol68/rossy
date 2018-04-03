package source

import (
	"github.com/yujiahaol68/rossy/app/database"
	"github.com/yujiahaol68/rossy/app/entity"
	"github.com/yujiahaol68/rossy/app/model/endpoint"
	"github.com/yujiahaol68/rossy/app/service"
)

const (
	unreadSourceList = `SELECT source_id, category_id, alias, name, COUNT(unread) AS count
	FROM (post LEFT JOIN source ON source_id = source.id)
	LEFT JOIN category ON category.id = category_id
	WHERE unread = 0 AND (post.delete_at IS NULL)
	GROUP BY source_id
	ORDER BY count;`
)

func FindOne(sourceID int64) *entity.Source {
	db := database.Conn()
	s := new(entity.Source)
	db.Id(sourceID).Get(s)
	return s
}

func FindByURL(url string) *entity.Source {
	db := database.Conn()
	s := new(entity.Source)
	db.Cols("id", "url").Where("source.url = ?", url).Get(s)
	return s
}

func Del(sourceID int64) {
	db := database.Conn()
	db.Id(sourceID).Delete(new(entity.Source))
}

func UnreadList() (*endpoint.UnreadSourceList, error) {
	db := database.Conn()
	rawRows, err := db.Query(unreadSourceList)
	if err != nil {
		return nil, err
	}

	u := make(endpoint.UnreadSourceList)

	for _, row := range rawRows {
		s := endpoint.Source{
			ID:          service.ToInt64(row["source_id"]),
			Category:    service.ToInt64(row["category_id"]),
			Alias:       string(row["alias"]),
			UnreadCount: service.ToInt64(row["count"]),
		}
		categoryName := string(row["name"])
		u[categoryName] = append(u[categoryName], s)
	}

	return &u, nil
}
