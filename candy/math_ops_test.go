package candy

import (
	"testing"
	"time"
)

func TestMathOps(t *testing.T) {
	if Max[int]() != 0 {
		t.Fatalf("expected zero")
	}
	if Max(1, 3, 2) != 3 {
		t.Fatalf("unexpected")
	}
	if Min[int]() != 0 {
		t.Fatalf("expected zero")
	}
	if Min(1, 3, 2) != 1 {
		t.Fatalf("unexpected")
	}

	if Sum(1, 2, 3) != 6 {
		t.Fatalf("unexpected")
	}
	if Average[int]() != 0 {
		t.Fatalf("unexpected")
	}
	if Average(2, 4) != 3 {
		t.Fatalf("unexpected")
	}
	if Average(time.Second, 3*time.Second) != 2*time.Second {
		t.Fatalf("unexpected")
	}

	if Abs(-3) != 3 {
		t.Fatalf("unexpected")
	}
	if Abs(3) != 3 {
		t.Fatalf("unexpected")
	}
	if Abs(-1.5) != 1.5 {
		t.Fatalf("unexpected")
	}
}
