package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"shield/common/utils"
)

func EncodePassword(password string) (salt string, passwordHash string) {
	salt = utils.NewUUID()
	h := sha256.New()

	h.Write([]byte(salt))
	h.Write([]byte(password))

	passwordHash = hex.EncodeToString(h.Sum(nil))

	return
}

func PasswordVerify(salt, passwordH, password string) bool {
	h := sha256.New()
	h.Write([]byte(salt))
	h.Write([]byte(password))

	return hex.EncodeToString(h.Sum(nil)) == passwordH
}
