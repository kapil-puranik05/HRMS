package models

type WeeklyReport struct {
	StudentID   uint   `json:"student_id"`
	Name        string `json:"name"`
	DaysPresent uint   `json:"days_present"`
}
