// Package web provides the web server
package web

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xavierxcn/chatgo/chatgo"
)

// Boot starts the web server
func Boot() {
	g := gin.Default()

	robot := chatgo.NewRobot().
		SetToken(os.Getenv("OPENAI_TOKEN")).
		SetName("Chatgo")

	go robot.Init()

	g.LoadHTMLGlob("./web/templates/*.tmpl")

	g.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	g.GET("/chatgo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"name": robot.Name()})
	})

	g.POST("/chatgo", func(c *gin.Context) {
		sentence := c.PostForm("sentence")
		sentence = strings.TrimSpace(sentence)

		if sentence != "" {
			robot.Tell(sentence)
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{"name": robot.Name(), "messages": robot.GetStringMessages()})
	})

	fmt.Println("Web server started at :8080, press Ctrl+C to stop")
	fmt.Println("access http://localhost:8080/chatgo to use Chatgo")

	if err := g.Run(":8080"); err != nil {
		panic(err)
	}
}
