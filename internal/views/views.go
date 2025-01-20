package views

import (
	"net/http"
	"strconv"

	"github.com/DevAthhh/todo/internal/database"
	"github.com/gin-gonic/gin"
)

func Views(router *gin.Engine) {
	router.GET("/", index)

	todo := router.Group("/t")
	{
		todo.POST("/create-node", create_node)
		todo.POST("/delete-node", delete_node)
		todo.POST("/update-node", update_node)
	}
}

func index(ctx *gin.Context) {
	todos := database.Select()

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"Todos": todos,
	})
}

func create_node(ctx *gin.Context) {
	title := ctx.PostForm("title")
	database.Insert(title)

	ctx.Redirect(302, "/")
}

func delete_node(ctx *gin.Context) {
	id := ctx.PostForm("id")
	database.Delete(id)

	ctx.Redirect(302, "/")
}

func update_node(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.PostForm("done"))
	todos := database.Select()
	for i := 0; i < len(todos); i++ {
		if todos[i].Id == id {
			database.Update(id, !todos[i].Done)
		}
	}

	ctx.Redirect(302, "/")
}
