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
