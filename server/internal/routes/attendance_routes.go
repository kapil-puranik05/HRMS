package routes

import (
	"server/internal/auth"
	"server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AttendanceRoutes(r *gin.Engine) {
	attendance := r.Group("/attendance")
	attendance.Use(auth.AuthMiddleware())
	{
		attendance.POST("/mark", controllers.MarkAttendance)
		attendance.GET("/:student_id", controllers.GetAttendance)
	}
}
