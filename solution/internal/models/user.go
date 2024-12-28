package models

type User struct {
	ID           int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Login        string `json:"login" gorm:"uniqueIndex;not null" validate:"required,min=3,max=50"`
	Email        string `json:"email" gorm:"uniqueIndex;not null" validate:"required,email"`
	Password     string `json:"password,omitempty" gorm:"-"`
	PasswordHash string `json:"-" gorm:"not null"`
	CountryCode  string `json:"countryCode" gorm:"size:3;not null" validate:"required,len=3"`
	IsPublic     bool   `json:"isPublic" gorm:"default:false"`
	Phone        string `json:"phone" gorm:"size:15" validate:"omitempty,min=10,max=15"`
	Image        string `json:"image" gorm:"size:255" validate:"omitempty,url"`
}
