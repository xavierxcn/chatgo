// Package web provides the web server
package web

import (
	"github.com/gin-gonic/gin"
)

func Boot() {
	g := gin.Default()

	g.LoadHTMLGlob("./web/templates/*.tmpl")

	g.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	g.GET("/chatgo", func(c *gin.Context) {

		c.HTML(200, "index.tmpl", gin.H{"name": "Chatgo"})

	})

	if err := g.Run(":8080"); err != nil {
		panic(err)
	}
}
