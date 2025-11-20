package services

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func DispatchReports() {
	c := cron.New()
	//Dispatching weekly reports
	c.AddFunc("0 23 * * 6", func() {
		SendWeeklyReports()
		fmt.Println("Sent weekly reports")
	})
	//Dispatching monthly reports
	c.AddFunc("0 23 28-31 * *", func() {
		now := time.Now()
		tomorrow := now.AddDate(0, 0, 1)

		if tomorrow.Day() == 1 {
			SendMonthlyReports()
			fmt.Println("Sent monthly reports")
		}
	})
	c.Start()
}

// Method for testing purposes -> Generates reports every 10s
func CronTest() {
	c := cron.New()
	c.AddFunc("@every 10s", func() {
		SendWeeklyReports()
		fmt.Println("Weekly Attendance Report")
	})
	c.AddFunc("@every 10s", func() {
		SendMonthlyReports()
		fmt.Println("Monthly Attendance Report")
	})
	c.Start()
}
