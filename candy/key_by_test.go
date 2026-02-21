package candy

import "testing"

type keyByPerson struct {
	ID      int
	ID8     int8
	ID16    int16
	ID32    int32
	ID64    int64
	UID     uint
	UID8    uint8
	UID16   uint16
	UID32   uint32
	UID64   uint64
	Score   float32
	Score64 float64
	Name    string
	Active  bool
}

func assertPanicsKeyBy(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatalf("expected panic")
		}
	}()
	fn()
}

func TestKeyBy(t *testing.T) {
	p1 := keyByPerson{
		ID: 1, ID8: 2, ID16: 3, ID32: 4, ID64: 5,
		UID: 6, UID8: 7, UID16: 8, UID32: 9, UID64: 10,
		Score: 1.5, Score64: 2.5, Name: "a", Active: true,
	}
	p2 := keyByPerson{
		ID: 2, ID8: 3, ID16: 4, ID32: 5, ID64: 6,
		UID: 7, UID8: 8, UID16: 9, UID32: 10, UID64: 11,
		Score: 2.5, Score64: 3.5, Name: "b", Active: false,
	}
	ss := []keyByPerson{p1, p2}

	if got := KeyByInt(ss, "ID"); got[1].Name != "a" || got[2].Name != "b" {
		t.Fatalf("unexpected: %v", got)
	}
	if got := KeyByInt8(ss, "ID8"); got[2].Name != "a" || got[3].Name != "b" {
		t.Fatalf("unexpected: %v", got)
	}
	if got := KeyByInt16(ss, "ID16"); got[3].Name != "a" || got[4].Name != "b" {
		t.Fatalf("unexpected: %v", got)
	}
	if got := KeyByInt32(ss, "ID32"); got[4].Name != "a" || got[5].Name != "b" {
		t.Fatalf("unexpected: %v", got)
	}
	if got := KeyByInt64(ss, "ID64"); got[5].Name != "a" || got[6].Name != "b" {
		t.Fatalf("unexpected: %v", got)
	}

	if got := KeyByUint(ss, "UID"); got[6].Name != "a" || got[7].Name != "b" {
		t.Fatalf("unexpected: %v", got)
	}
	if got := KeyByUint8(ss, "UID8"); got[7].Name != "a" || got[8].Name != "b" {
		t.Fatalf("unexpected: %v", got)
	}
	if got := KeyByUint16(ss, "UID16"); got[8].Name != "a" || got[9].Name != "b" {
		t.Fatalf("unexpected: %v", got)
	}
	if got := KeyByUint32(ss, "UID32"); got[9].Name != "a" || got[10].Name != "b" {
		t.Fatalf("unexpected: %v", got)
	}
	if got := KeyByUint64(ss, "UID64"); got[10].Name != "a" || got[11].Name != "b" {
		t.Fatalf("unexpected: %v", got)
	}

	if got := KeyByFloat32(ss, "Score"); got[1.5].Name != "a" || got[2.5].Name != "b" {
		t.Fatalf("unexpected: %v", got)
	}
	if got := KeyByFloat64(ss, "Score64"); got[2.5].Name != "a" || got[3.5].Name != "b" {
		t.Fatalf("unexpected: %v", got)
	}
	if got := KeyByString(ss, "Name"); got["a"].ID != 1 || got["b"].ID != 2 {
		t.Fatalf("unexpected: %v", got)
	}
	if got := KeyByBool(ss, "Active"); got[true].Name != "a" || got[false].Name != "b" {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestKeyBy_Pointers(t *testing.T) {
	p1 := &keyByPerson{ID: 1, Name: "a"}
	p2 := &keyByPerson{ID: 2, Name: "b"}
	got := KeyByInt([]*keyByPerson{p1, p2}, "ID")
	if got[1].Name != "a" || got[2].Name != "b" {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestKeyBy_Panics(t *testing.T) {
	t.Run("field not found", func(t *testing.T) {
		assertPanicsKeyBy(t, func() {
			KeyByInt([]keyByPerson{{ID: 1}}, "Nope")
		})
	})

	t.Run("wrong field type", func(t *testing.T) {
		assertPanicsKeyBy(t, func() {
			KeyByInt([]keyByPerson{{ID: 1}}, "Name")
		})
	})

	t.Run("non-struct element", func(t *testing.T) {
		assertPanicsKeyBy(t, func() {
			KeyByInt([]int{1, 2, 3}, "X")
		})
	})

	t.Run("pointer to non-struct", func(t *testing.T) {
		assertPanicsKeyBy(t, func() {
			KeyByInt([]*int{new(int)}, "X")
		})
	})
}
