package main

import (
	"cron/databases"
	"cron/initializers"
	"cron/services"
)

func main() {
	initializers.SetEnvironment()
	databases.ConnectDB()
	services.InitiateMailService()
	go services.DispatchReports()
	//Testing the cron setup with a cron job that runs every minute,
	//performing the same task as the DispatchReports() Method

	//go services.CronTest() //Remove the comment from this method to test the working of weekly and monthly report generation(Every 10s)
	select {}
}
