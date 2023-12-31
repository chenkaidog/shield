package errs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	var err error = nil

	err = &bizError{}

	_, ok := err.(Error)
	assert.True(t, ok)

	_, ok = err.(*bizError)
	assert.True(t, ok)
}
