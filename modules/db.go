package modules

import (
  "github.com/google/uuid"
  "golang.org/x/crypto/bcrypt"
)

func GenerateUUID() (newUuid string) {
  newUuidObject, _ := uuid.NewUUID()
  newUuid = newUuidObject.String()
  return newUuid
}

func Hashization(text string) (hashText string) {
  hashText, _ = bcrypt.GenerateFromPassword([]byte(text),12)
  return hashText
}
