package api

import (
	"net/http"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	wd, _ := os.Getwd()
	fmt.Print(wd)
	engine.Static("/css", "src/css")
	engine.Static("/js", "src/js")

	engine.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
        engine.GET("/register", func(c *gin.Context) {
                c.HTML(http.StatusOK, "register.html", gin.H{})
        })
}
