package modules

import (
  "os"
  
  "gopkg.in/go-ini/ini.v1"
)

func GetServerConfig() (serverConfig ServerConfig) {
  cfg, err := ini.Load("server.conf")
  if err != nil {
    os.Exit(-1)
  }
  serverConfig = ConfigList{
	  Port:      cfg.Section("server").Key("port").MustInt(8080),
	  SSL:      cfg.Section("server").Key("SSL").MustBool(false),
	  ServerCertificate:      cfg.Section("server").Key("port").String(),
    PrivateKey: cfg.Section("server").Key("port").String(),
    }
  return serverConfig
}

func GetDBConfig() () {
  
  
}
