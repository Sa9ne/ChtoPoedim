package models

type Users struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
