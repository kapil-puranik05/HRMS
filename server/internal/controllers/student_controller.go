package controllers

import (
	"net/http"
	"server/internal/models"
	"server/internal/services"
	"server/internal/views"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateStudent godoc
// @Summary Create a new student
// @Description Creates a new student in the database
// @Tags Students
// @Accept json
// @Produce json
// @Param student body views.CreateStudentRequest true "Student data"
// @Success 201 {object} models.Student
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Router /students/ [post]
func CreateStudent(c *gin.Context) {
	var req views.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student := models.Student{
		Name:       req.Name,
		Email:      req.Email,
		Department: req.Department,
		CreatedAt:  time.Now(),
	}
	created, err := services.CreateStudent(student)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create student"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "student created successfully", "data": created})
}

// GetStudent godoc
// @Summary Get a student by ID
// @Description Loads a single student using ID
// @Tags Students
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} models.Student
// @Failure 404 {object} map[string]string
// @Router /students/{id} [get]
func GetStudent(c *gin.Context) {
	idStr := c.Param("id")
	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	student, err := services.GetStudent(uint(idUint64))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "student data loaded successfully", "data": student})
}

// GetStudents godoc
// @Summary Get all students
// @Description Returns a list of all students
// @Tags Students
// @Produce json
// @Success 200 {array} models.Student
// @Router /students/ [get]
func GetStudents(c *gin.Context) {
	students, err := services.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load students"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "student data loaded successfully", "data": students})
}

// UpdateStudent godoc
// @Summary Update existing student
// @Description Updates student details by ID
// @Tags Students
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param student body views.UpdateStudentRequest true "Updated data"
// @Success 200 {object} models.Student
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /students/{id} [put]
func UpdateStudent(c *gin.Context) {
	idStr := c.Param("id")
	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	student, err := services.GetStudent(uint(idUint64))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}
	var req views.UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Name != nil {
		student.Name = *req.Name
	}
	if req.Email != nil {
		student.Email = *req.Email
	}
	if req.Department != nil {
		student.Department = *req.Department
	}
	if err := services.UpdateStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update student data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "student updated successfully", "data": student})
}

// DeleteStudent godoc
// @Summary Delete a student
// @Description Deletes a student by ID
// @Tags Students
// @Param id path int true "Student ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /students/{id} [delete]
func DeleteStudent(c *gin.Context) {
	idStr := c.Param("id")
	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := services.DeleteStudentByID(uint(idUint64)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "student deleted successfully"})
}
