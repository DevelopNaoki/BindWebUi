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

func HashComparison(hashText string, text string) (agreement bool) {
  err := bcrypt.CompareHashAndPassword([]byte(hashText), []byte(text))
  if err != nil {
     agreement = false
  } else {
    agreement = true
  }
  return agreement
}

