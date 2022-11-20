package utils

import (
	"strings"

	"github.com/google/uuid"
)

func UUIDNoHyphen() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
