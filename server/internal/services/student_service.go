package services

import (
	"errors"
	"fmt"

	"server/internal/databases"
	"server/internal/models"

	"gorm.io/gorm"
)

func GetStudent(id uint) (models.Student, error) {
	var student models.Student
	if err := databases.DB.First(&student, id).Error; err != nil {
		return student, fmt.Errorf("error occured while fetching student data: %v", err.Error())
	}
	return student, nil
}

func GetStudents() ([]models.Student, error) {
	var students []models.Student
	if err := databases.DB.Find(&students).Error; err != nil {
		return nil, fmt.Errorf("error occured while fetching students: %v", err.Error())
	}
	return students, nil
}

func CreateStudent(student models.Student) (models.Student, error) {
	var existing models.Student
	result := databases.DB.Where("email = ?", student.Email).First(&existing)
	if result.Error == nil && existing.ID != 0 {
		return existing, fmt.Errorf("student already exists with email %s", student.Email)
	}
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) && result.Error != nil {
		return student, fmt.Errorf("error occured while finding student by email: %v", result.Error.Error())
	}
	if err := databases.DB.Create(&student).Error; err != nil {
		return student, fmt.Errorf("error occured while creating student: %v", err.Error())
	}
	return student, nil
}

func UpdateStudent(student *models.Student) error {
	if err := databases.DB.Save(student).Error; err != nil {
		return fmt.Errorf("error occured while updating student: %v", err.Error())
	}
	return nil
}

func DeleteStudentByID(id uint) error {
	var student models.Student
	if err := databases.DB.First(&student, id).Error; err != nil {
		return fmt.Errorf("error occured while fetching student data: %v", err.Error())
	}
	if err := databases.DB.Delete(&student).Error; err != nil {
		return fmt.Errorf("error occured while deleting student: %v", err.Error())
	}
	return nil
}
