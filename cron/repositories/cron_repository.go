package repositories

import (
	"cron/databases"
	"cron/models"
	"fmt"
)

func GenerateWeeklyReport(studentID uint) (models.WeeklyReport, error) {
	var report models.WeeklyReport
	result := databases.DB.Raw(`
		select attendances.student_id as student_id, name, count(*) as days_present
		from attendances 
		inner join students 
		on attendances.student_id = students.id
		where attendances.student_id = ? 
		and date between date_sub(curdate(), interval 1 week) and curdate()
		group by attendances.student_id;
	`, studentID).Scan(&report)
	if result.Error != nil {
		return report, fmt.Errorf("error occured while generating weekly report for student %d: %v", studentID, result.Error)
	}
	return report, nil
}

func GenerateMonthlyReport(studentID uint) (models.MonthlyReport, error) {
	var report models.MonthlyReport
	result := databases.DB.Raw(`
		select attendances.student_id as student_id, name, count(attendances.id) as days_present
		from attendances
		inner join students
		on attendances.student_id = students.id where attendances.student_id = ?
		and month(date) = month(curdate()) and year(date) = year(curdate())
		group by attendances.student_id;
	`, studentID).Scan(&report)
	if result.Error != nil {
		return report, fmt.Errorf("error occured while generating monthly report for student %d: %v", studentID, result.Error)
	}
	return report, nil
}

//A functional endpoint method for persisting the weekly and monthly reports in the database can also be added here
