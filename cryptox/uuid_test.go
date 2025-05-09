package cryptox

import (
	"testing"
)

func TestUUID(t *testing.T) {
	result := UUID()
	if len(result) != 32 {
		t.Errorf("expected 32 characters, got %d", len(result))
	}
}
