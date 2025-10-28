package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;<-:create" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Username  string    `gorm:"not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	Tasks     []Task    `gorm:"foreignKey:ID" json:"tasks"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
