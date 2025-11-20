package models

type MonthlyReport struct {
	StudentID   uint   `json:"student_id"`
	Name        string `json:"name"`
	DaysPresent uint   `json:"days_present"`
}
