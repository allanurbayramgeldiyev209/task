package routes

import (
	"task/controllers"
	"task/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	routes := gin.Default()

	routes.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "RefreshToken", "Authorization"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		// MaxAge:           12 * time.Hour,
	}))

	api := routes.Group("/api")
	{
		api.POST("resgiter", controllers.RegisterUser)
		api.POST("login", controllers.LoginUser)
		api.GET("lessons", middlewares.CheckUser(), controllers.GetLessons)
		api.PUT("lessons", middlewares.CheckUser(), controllers.UpdateLessons)

	}

	return routes
}
