package models

import "encoding/json"

type UserPublic struct {
	Username          string          `json:"username" gorm:"primaryKey;unique;not null"`
	Name              *string         `json:"name"`
	ProfessionalTitle *string         `json:"professional_title"`
	CVJson            json.RawMessage `json:"cv_json" gorm:"type:jsonb" swaggerignore:"true"`
	CompanyName       *string         `json:"company_name"`
	Location          *string         `json:"location"`
	Description       *string         `json:"description"`
	Image             []Image         `gorm:"foreignKey:Username;references:Username"`
}

func (UserPublic) TableName() string {
	return "User"
}

type User struct {
	Username          string          `json:"username" gorm:"primaryKey;unique;not null"`
	Name              *string         `json:"name"`
	Password          string          `json:"password" gorm:"not null"`
	ProfessionalTitle *string         `json:"professional_title"`
	CVJson            json.RawMessage `json:"cv_json" gorm:"type:jsonb" swaggerignore:"true"`
	CompanyName       *string         `json:"company_name"`
	Location          *string         `json:"location"`
	Description       *string         `json:"description"`
}

func (User) TableName() string {
	return "User"
}
