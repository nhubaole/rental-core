package models

import "time"

type User struct {
	ID          int        `gorm:"primaryKey;autoIncrement"`
	PhoneNumber string     `gorm:"type:text;not null" json:"phoneNumber"`
	FullName    string     `gorm:"type:text;not null" json:"fullName"`
	Address     string     `gorm:"type:text;not null" json:"address"`
	Password    string     `gorm:"type:text;not null" json:"password"`
	Gender      *int       `gorm:"type:text"`
	Birthday    *time.Time `gorm:"type:text"`
	Role        int        `gorm:"type:integer;not null" json:"role"`
	CreatedAt   time.Time
}
