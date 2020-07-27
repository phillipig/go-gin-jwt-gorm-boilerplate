package Routes

import (
	"go-api/Config"
	"go-api/Controllers"
	"go-api/Middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func InitGin() {
	authMiddleware, err := Middleware.Auth()

	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Middleware.Cors())

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Route not found"})
	})

	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/user-api")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("user", Controllers.GetUsers)
		auth.POST("user", Controllers.CreateUser)
		auth.GET("user/:id", Controllers.GetUserByID)
		auth.PUT("user/:id", Controllers.UpdateUser)
		auth.DELETE("user/:id", Controllers.DeleteUser)
	}

	r.Run(":" + Config.GetEnvKey("API_GIN_PORT"))
}
