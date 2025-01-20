package handlers

import (
	"github.com/DevAthhh/todo/internal/views"
	"github.com/gin-gonic/gin"
)

func Handle() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("internal/templates/*")
	router.Static("static/", "internal/static/")

	views.Views(router)

	return router
}
