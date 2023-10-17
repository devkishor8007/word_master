package models

import (
	"time"
)

type Article struct {
	ArticleID       uint      `gorm:"primaryKey" json:"article_id"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	PublicationDate time.Time `json:"publication_date"`
	AuthorID        uint      `json:"author_id"`
	User            User      `json:"user" gorm:"foreignKey:AuthorID"`
	CategoryID      uint      `json:"category_id"`
	Category        Category  `json:"category" gorm:"foreignKey:CategoryID"`
	Comments        []Comment `gorm:"foreignKey:SArticleID" json:"comments"`
}
