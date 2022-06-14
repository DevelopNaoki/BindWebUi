package modules

type User struct {
    Id int
    Uuid string
    Username string
    Password string
}

type Role steuct {
    Id int
    Uuid string
    Role string
}

type UserRole struct {
    Id int
    UserUuid string
    RoleUuid string
}
