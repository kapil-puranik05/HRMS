package controllers

import (
	"net/http"
	"server/internal/databases"
	"server/internal/models"
	"server/internal/views"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAttendance(c *gin.Context) {
	id := c.Param("student_id")
	var attendance []models.Attendance
	result := databases.DB.Where("student_id = ?", id).Find(&attendance)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "student not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "student attendance data found",
		"data":    attendance,
	})
}

// MarkAttendance godoc
// @Summary Mark attendance for a student
// @Description Marks today's attendance for a student. Fails if already marked or student does not exist.
// @Tags Attendance
// @Accept json
// @Produce json
// @Param attendance body views.AttendanceRequest true "Attendance data"
// @Success 201 {object} models.Attendance
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /attendance/mark [post]
func MarkAttendance(c *gin.Context) {
	var req views.AttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	attendance := models.Attendance{
		StudentId: req.StudentID,
		Status:    "PRESENT",
		Date:      time.Now(),
	}
	result := databases.DB.Create(&attendance)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "foreign key constraint") {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "student not found",
			})
			return
		}
		if strings.Contains(result.Error.Error(), "Duplicate entry") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "attendance already marked for today",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "attendance marked successfully",
		"data":    attendance,
	})
}
