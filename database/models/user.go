package models

type User struct {
	Username string `json:"username" gorm:"primary_key;unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

func (User) TableName() string {
	return "User"
}
