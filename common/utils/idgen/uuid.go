package idgen

import (
	"encoding/hex"

	"github.com/google/uuid"
)

func NewUUID() string {
	binId, _ := uuid.New().MarshalBinary()
	return hex.EncodeToString(binId)
}
