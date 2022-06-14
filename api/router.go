package api

import (
	"net/http"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	p, _ := os.Getwd()
	fmt.Print(p)
	engine.Static("/css", "src/css")
	engine.Static("/js", "src/js")

	engine.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
}
