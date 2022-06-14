package api

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	engine := gin.Default()
	engine.LoadHTMLGlob("src/templates/*.html")

	Router(engine)

	engine.Run(":3000")
}
