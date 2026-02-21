package candy

import "testing"

type pluckPerson struct {
	Name string
	Age  int
	City string
	Tags []string
}

func assertPanicsPluck(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatalf("expected panic")
		}
	}()
	fn()
}

func TestPluck_Generic(t *testing.T) {
	if got := Pluck([]pluckPerson{}, func(p pluckPerson) string { return p.Name }); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}

	got := Pluck([]pluckPerson{{Name: "a"}, {Name: "b"}}, func(p pluckPerson) string { return p.Name })
	if len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestPluckPtr(t *testing.T) {
	p1 := &pluckPerson{Name: "a"}
	got := PluckPtr([]*pluckPerson{p1, nil}, func(p *pluckPerson) string { return p.Name }, "x")
	if len(got) != 2 || got[0] != "a" || got[1] != "x" {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestPluckUnique(t *testing.T) {
	got := PluckUnique([]pluckPerson{{City: "x"}, {City: "x"}, {City: "y"}}, func(p pluckPerson) string { return p.City })
	if len(got) != 2 {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestPluckMapAndGroupBy(t *testing.T) {
	items := []pluckPerson{{Name: "a", City: "x"}, {Name: "b", City: "x"}, {Name: "c", City: "y"}}

	m := PluckMap(items, func(p pluckPerson) string { return p.Name }, func(p pluckPerson) string { return p.City })
	if m["a"] != "x" || m["c"] != "y" {
		t.Fatalf("unexpected: %v", m)
	}

	g := PluckGroupBy(items, func(p pluckPerson) string { return p.City })
	if len(g["x"]) != 2 || len(g["y"]) != 1 {
		t.Fatalf("unexpected: %v", g)
	}
}

func TestPluck_Reflect(t *testing.T) {
	items := []pluckPerson{{Name: "a", Age: 1, Tags: []string{"t1"}}, {Name: "b", Age: 2, Tags: []string{"t2"}}}

	if got := PluckString(items, "Name"); len(got) != 2 || got[0] != "a" || got[1] != "b" {
		t.Fatalf("unexpected: %v", got)
	}
	if got := PluckInt(items, "Age"); len(got) != 2 || got[0] != 1 || got[1] != 2 {
		t.Fatalf("unexpected: %v", got)
	}
	if got := PluckStringSlice(items, "Tags"); len(got) != 2 || got[0][0] != "t1" || got[1][0] != "t2" {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestPluck_ReflectPanics(t *testing.T) {
	assertPanicsPluck(t, func() {
		PluckString([]pluckPerson{{Name: "a"}}, "Nope")
	})

	assertPanicsPluck(t, func() {
		PluckString([]int{1, 2, 3}, "X")
	})

	assertPanicsPluck(t, func() {
		PluckString("x", "Y")
	})
}
