package models

type User struct {
	UserID   uint   `gorm:"primaryKey" json:"user_id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}