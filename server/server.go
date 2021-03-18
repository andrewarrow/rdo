package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rdo/controllers"
)

func RoutesSetup(router *gin.Engine) {

	router.Static("/static", "static")
	router.GET("/", controllers.WelcomeIndex)
	router.GET("/privacy", controllers.LegalPrivacy)
	router.GET("/terms", controllers.LegalTerms)
	api := router.Group("/api")
	api.GET("/version", controllers.ApiVersion)
	users := router.Group("/users")
	users.GET("/", controllers.UsersIndex)
	users.GET("/:id", controllers.UsersShow)

	user := router.Group("/user")
	user.GET("/new", controllers.UsersNew)
	user.POST("/create", controllers.UsersCreate)

	sessions := router.Group("/sessions")
	sessions.GET("/new", controllers.SessionsNew)
	sessions.POST("/", controllers.SessionsCreate)
	sessions.POST("/destroy", controllers.SessionsDestroy)

	admin := router.Group("/admin")
	users = admin.Group("/users")
	users.GET("/", controllers.AdminUsersIndex)
	user = admin.Group("/user")
	user.GET("/:id", controllers.AdminUsersShow)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	AddTemplates(router)
}
