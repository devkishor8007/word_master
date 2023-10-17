package models

import (
	"time"
)

type Comment struct {
	CommentID   uint      `gorm:"primaryKey" json:"comment_id"`
	Text        string    `json:"text"`
	CommentDate time.Time `json:"comment_date"`
	AuthorID    uint      `json:"author_id"`
	User        User      `json:"user" gorm:"foreignKey:AuthorID"`
	SArticleID  uint      `json:"article_id"`
	Article     Article   `json:"article" gorm:"foreignKey:SArticleID"`
}
