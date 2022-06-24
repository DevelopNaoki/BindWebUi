package modules

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)
    
type User struct {
    gorm.Model
    Uuid string
    Username string
    Password string
}

type Role struct {
    gorm.Model
    Uuid string
    Role string
}

type UserRole struct {
    gorm.Model
    UserUuid string
    RoleUuid string
}

type ServerConfig struct {
    Port string
    SSL bool
    ServerCertificate string
    PrivateKey string
}

type DBConfig struct {
    MySQLServer string
    Database string
    Username string
    Password string
}
