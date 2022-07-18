package models

import "time"

type Post struct {
	ID          int64     `json:"id" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id"`
	status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title"`
	Context     string    `json:"context" db:"context"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}
