package router

import (
	"squirtle/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", apis.Users)
	router.POST("/user", apis.Add)
	router.PUT("/user/:id", apis.Update)
	router.DELETE("/user/:id", apis.Delete)

	return router
}
