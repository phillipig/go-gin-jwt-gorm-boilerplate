package routes

import (
	"go-api/config"
	"go-api/controllers"
	"go-api/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func NewGin() {
	authMiddleware, err := middleware.Auth()
	env := config.NewEnv()

	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	if env.GetEnvKeyBool("GIN_RELEASEMODE") {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.POST("/login", authMiddleware.LoginHandler)

	//User
	user := r.Group("/user")
	user.Use(authMiddleware.MiddlewareFunc())
	{
		controller := controllers.NewUserController()
		user.GET("", controller.ReadAll)
		user.POST("", controller.Create)
		user.GET("/:id", controller.ReadByID)
		user.PUT("/:id", controller.Update)
		user.DELETE("/:id", controller.Delete)
	}

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Route not found"})
	})

	r.Run(":" + env.GetEnvKey("API_GIN_PORT"))
}
