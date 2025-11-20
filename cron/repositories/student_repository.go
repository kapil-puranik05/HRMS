package repositories

import (
	"cron/databases"
	"cron/models"
	"fmt"
)

func GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	result := databases.DB.Find(&students)
	if result.Error != nil {
		return nil, fmt.Errorf("error occured while fetching students: %v", result.Error.Error())
	}
	return students, nil
}
