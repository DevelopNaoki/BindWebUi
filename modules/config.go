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

  serverConfig.Port              = cfg.Section("server").Key("port").MustInt(8080)
  serverConfig.SSL               = cfg.Section("server").Key("SSL").MustBool(false)
  serverConfig.ServerCertificate = cfg.Section("server").Key("server_certificate").String()
  serverConfig.PrivateKey        = cfg.Section("server").Key("private_key").String()

  return serverConfig
}

func GetDBConfig() (dbConfig DBConfig) {
  cfg, err := ini.Load("server.conf")
  if err != nil {
    os.Exit(-1)
  }

  dbConfig.MySQKServer = cfg.Section("db").Key("mysql_ip").String()
  dbConfig.Database    = cfg.Section("db").Key("database").String()
  dbConfig.Username    = cfg.Section("db").Key("username").String()
  dbConfig.Password    = cfg.Section("db").Key("password").String()

  return dbConfig
}

func FileExist (filename string) (exist bool) {

}
