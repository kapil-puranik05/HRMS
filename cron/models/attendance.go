package models

import "time"

type Attendance struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	StudentId uint      `json:"student_id" gorm:"uniqueIndex:attendance_unique"`
	Student   Student   `gorm:"constraint:OnDelete:CASCADE"`
	Date      time.Time `json:"date" gorm:"type:date;uniqueIndex:attendance_unique"`
	Status    string    `json:"status"`
}
