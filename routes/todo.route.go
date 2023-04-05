package routes

import (
	"gin-todo-list/models"
	"gin-todo-list/repositories"
	"log"

	"github.com/gin-gonic/gin"
)

func todoRouter(router *gin.RouterGroup) {
	// Ping test
	router.GET("/", func(c *gin.Context) {
		_, result := repositories.Todo.Find()
		c.JSON(200, result)
		// c.String(http.StatusOK, "pong")
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

	router.POST("/", func(c *gin.Context) {
		var todo models.Todo

		if err := c.BindJSON(&todo); err != nil {
			log.Fatalln(err)
		}

		repositories.Todo.Create(todo)
		c.Status(200)
	})
}
