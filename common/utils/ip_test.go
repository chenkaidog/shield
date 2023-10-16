package utils

import (
	"testing"
)

func TestIPv4(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		t.Log(IPv4())
	})

	t.Run("hex", func(t *testing.T) {
		t.Log(IPv4Hex())
	})

	t.Run("int", func(t *testing.T) {
		t.Logf("%d", IPv4Int())
	})
}
