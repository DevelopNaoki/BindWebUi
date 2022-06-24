package api

import (
	"github.com/gin-gonic/gin"
	
	"github.com/DevelopNaoki/pcawui/modules"
)

func Start() {
	serverConfig := modules.GetServerConfig()

	engine := gin.Default()
	engine.LoadHTMLGlob("src/templates/*.html")

	Router(engine)

	if serverConfig.SSL {
		engin.RunTLS(":"+serverConfig.Port, serverConfig.ServerCertificate, serverConfig.PrivateKey)
	} else {
		engine.Run(":"+serverConfig.Port)
	}
}
