package models

type Dishes struct {
	ID          uint   `gorm:"PrimaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Time        int    `json:"time"`
}
