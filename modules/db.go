package modules

import (
  "github.com/google/uuid"
)

func GenerateUUID() (newUuid string) {
  newUuidObject, _ := uuid.NewUUID()
  newUuid = newUuidObject.String()
  return newUuid
}
