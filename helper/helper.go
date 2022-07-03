package helper

import (
	"github.com/google/uuid"
)

func UUIDMaker() string {
	id := uuid.New()
	return id.String()
}
