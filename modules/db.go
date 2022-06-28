package modules

import (
  "github.com/google/uuid"
  "golang.org/x/crypto/bcrypt"
)

func GenerateUUID() (new_uuid string) {
  new_uuid_object, _ := uuid.NewUUID()
  new_uuid = new_uuid_object.String()
  return new_uuid
}

func Hashization(plain_text string) (hash_text string) {
  hash_text, _ = bcrypt.GenerateFromPassword([]byte(plain_text),12)
  return hash_text
}

func HashComparison(hash_text string, plain_text string) (agreement bool) {
  err := bcrypt.CompareHashAndPassword([]byte(hash_text), []byte(plain_text))
  if err != nil {
     agreement = false
  } else {
    agreement = true
  }
  return agreement
}

