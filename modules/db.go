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

func Hashization(plainText string) (hashText string) {
  hashText, _ = bcrypt.GenerateFromPassword([]byte(plainText),12)
  return hashText
}

func HashComparison(hashText string, plainText string) (agreement bool) {
  err := bcrypt.CompareHashAndPassword([]byte(hashText), []byte(plainText))
  if err != nil {
     agreement = false
  } else {
    agreement = true
  }
  return agreement
}

