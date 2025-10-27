package model

import "time"

type Task struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Status    bool      `gorm:"default:false" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTask(Name string) *Task {
	return &Task{Name: Name}
}
