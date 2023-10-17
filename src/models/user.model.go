package models

type User struct {
	UserID   uint      `gorm:"primaryKey" json:"user_id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Articles []Article `gorm:"foreignKey:AuthorID" json:"articles"`
	Comments []Comment `gorm:"foreignKey:AuthorID" json:"comments"`
}
