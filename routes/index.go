package routes

import "github.com/gin-gonic/gin"

func Setup(app *gin.Engine) {
	todoRouter(app.Group("/todo"))
	todoListRouter(app.Group("/todo-list"))
}
