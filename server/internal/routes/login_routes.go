package routes

import (
	"server/internal/auth"

	"github.com/gin-gonic/gin"
)

func LoginRoutes(r *gin.Engine) {
	public := r.Group("/public")
	{
		public.POST("/login", auth.Login)
	}
}
