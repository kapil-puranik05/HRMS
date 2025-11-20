package models

import "time"

type Student struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Email      string    `json:"email" gorm:"type:varchar(255);unique"`
	Department string    `json:"department"`
	CreatedAt  time.Time `json:"created_at"`
}
