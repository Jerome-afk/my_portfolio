package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("*.html")
	router.Static("/static", "./static")

	router.GET("/", func (c *gin.Context) {
		c.Redirect(http.StatusFound, "/home")
	})

	router.GET("/home", func (c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", nil)
	})

	return router
}