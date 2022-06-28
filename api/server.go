package api

import (
	"github.com/gin-gonic/gin"
	
	"github.com/DevelopNaoki/pcawui/modules"
)

func Start() {
	serverConfig := modules.FetchServerConfig()

	engine := gin.Default()
	engine.LoadHTMLGlob("src/templates/*.html")

	Router(engine)

	if serverConfig.SSLEnabled {
		engin.RunTLS(":"+serverConfig.Port, serverConfig.ServerCertificate, serverConfig.PrivateKey)
	} else {
		engine.Run(":"+serverConfig.Port)
	}
}
