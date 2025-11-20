package routes

import (
	"server/internal/auth"
	"server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.Engine) {
	students := r.Group("/students")
	students.Use(auth.AuthMiddleware())
	{
		students.POST("/", controllers.CreateStudent)
		students.GET("/", controllers.GetStudents)
		students.GET("/:id", controllers.GetStudent)
		students.PUT("/:id", controllers.UpdateStudent)
		students.DELETE("/:id", controllers.DeleteStudent)
	}
}
