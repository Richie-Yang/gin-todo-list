package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func todoListRouter(router *gin.RouterGroup) {
	// Ping test
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	router.GET("/:id", func(c *gin.Context) {
		// user := c.Params.ByName("name")
		// value, ok := datasource.DB[user]
		// if ok {
		// 	c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		// } else {
		// 	c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		// }
	})
}
