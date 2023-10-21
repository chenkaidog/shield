package idgen

import "testing"

func TestIdGen(t *testing.T) {
	idGen := NewIDGenerator(10)
	t.Run("trace id", func(t *testing.T) {
		t.Log(idGen.NewTraceID())
	})

	t.Run("log id", func(t *testing.T) {
		t.Log(idGen.NewLogID())
	})

	t.Run("span id", func(t *testing.T) {
		t.Log(idGen.NewSpanID(""))

		t.Log(idGen.NewSpanID("psanid"))
	})
}
