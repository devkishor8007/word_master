package models

import (
    "time"
)

type Comment struct {
    CommentID  uint      `gorm:"primaryKey" json:"comment_id"`
    Text       string    `json:"text"`
    CommentDate time.Time `json:"comment_date"`
    UserID     uint      `json:"user_id"`
    ArticleID  uint      `json:"article_id"`
}