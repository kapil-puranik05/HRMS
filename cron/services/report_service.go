package services

import (
	"cron/models"
	"cron/repositories"
	"fmt"
	"log"
)

func PrettyWeeklyReport(report *models.WeeklyReport) string {
	return fmt.Sprintf(
		"Weekly attendance report for %s.<br><br>Total days present: %d<br><br>This is an automated email. Please do not reply.",
		report.Name,
		report.DaysPresent,
	)
}

func PrettyMonthlyReport(report *models.MonthlyReport) string {
	return fmt.Sprintf(
		"Monthly attendance report for %s.<br><br>Total days present: %d<br><br>This is an automated email. Please do not reply.",
		report.Name,
		report.DaysPresent,
	)
}

func SendWeeklyReports() {
	students, err := repositories.GetAllStudents()
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, s := range students {
		report, err := repositories.GenerateWeeklyReport(s.ID)
		if err != nil {
			log.Println(err.Error())
			return
		}
		result := PrettyWeeklyReport(&report)
		Mailer.SendMail(
			s.Email,
			"Your Weekly Attendance Report",
			result,
		)
		fmt.Println(result)
	}
}

func SendMonthlyReports() {
	students, err := repositories.GetAllStudents()
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, s := range students {
		report, err := repositories.GenerateMonthlyReport(s.ID)
		if err != nil {
			log.Println(err.Error())
			return
		}
		result := PrettyMonthlyReport(&report)
		Mailer.SendMail(
			s.Email,
			"Your Monthly Attendance Report",
			result,
		)
		fmt.Println(result)
	}
}
