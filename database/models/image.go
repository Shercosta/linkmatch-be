package models

type Image struct {
	ID       string `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ImageURL string `json:"image_url"`
	Username string `json:"username"`
	Position int    `json:"position"`
	User     User   `gorm:"foreignKey:Username;references:Username"`
}

func (Image) TableName() string {
	return "Image"
}
