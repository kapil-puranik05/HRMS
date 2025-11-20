// @title HRMS API
// @version 1.0
// @description API documentation for HRMS System built in Go (Gin)
// @host localhost:8080
// @BasePath /
package main

import (
	"fmt"
	"log"
	"os"
	_ "server/docs"
	"server/initializers"
	"server/internal/databases"
	"server/internal/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.SetEnvironment()
	databases.ConnectDB()
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.StudentRoutes(r)
	routes.AttendanceRoutes(r)
	routes.LoginRoutes(r)
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal("Failed to start the application")
	} else {
		fmt.Println("Server Started on Port 8080")
	}
}
