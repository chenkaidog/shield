package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodePassword(t *testing.T) {
	salt1, pwd1 := EncodePassword("123")
	salt2, pwd2 := EncodePassword("123")

	t.Log(salt1, salt2)
	t.Log(pwd1, pwd2)
	if salt1 == salt2 {
		assert.Equal(t, pwd1, pwd2)
	} else {
		assert.NotEqual(t, pwd1, pwd2)
	}

	assert.True(t, PasswordVerify(salt1, pwd1, "123"))
	assert.True(t, PasswordVerify(salt2, pwd2, "123"))

	assert.False(t, PasswordVerify(salt1, pwd2, "123"))
	assert.False(t, PasswordVerify(salt1, pwd1, "1234"))
}
