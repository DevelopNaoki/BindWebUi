package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	p, _ := os.Getwd()
	fmt.Print(p)
	engine.Static("/css", "src/css")
	engine.Static("/js", "src/js")
}
