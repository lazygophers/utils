package candy

import "testing"

func TestEqual(t *testing.T) {
	if !Equal(1, 1) {
		t.Fatalf("expected true")
	}
	if Equal(1, 2) {
		t.Fatalf("expected false")
	}
}
