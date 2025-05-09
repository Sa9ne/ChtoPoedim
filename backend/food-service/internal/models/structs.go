package models

type Dishes struct {
	ID          uint   `gorm:"PrimaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Img         string `json:"image"`
	Time        int    `json:"time"`
	Vegan       bool   `json:"vegan"`
	Spicy       bool   `json:"spicy"`
	Hearty      bool   `json:"hearty"`
}

type UserChoice struct {
	Stage  string `json:"stage"`
	Choice bool   `json:"choice"`
}

type Users struct {
	ID        uint     `json:"id" gorm:"primaryKey"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Favorites []Dishes `gorm:"many2many:user_favorites"`
}
