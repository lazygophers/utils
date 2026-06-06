package language

import (
	"sync"
	"testing"
)

func TestDefault(t *testing.T) {
	if Default().String() != "en" {
		t.Errorf("Default() = %q, want %q", Default().String(), "en")
	}
}

func TestSetDefault(t *testing.T) {
	orig := Default()
	defer SetDefault(orig)

	SetDefault(Make("zh-CN"))
	if Default().String() != "zh-CN" {
		t.Errorf("after SetDefault(zh-CN): got %q", Default().String())
	}

	SetDefault(nil)
	if Default() != nil {
		t.Errorf("SetDefault(nil): got %v, want nil", Default())
	}
}

func TestGet_NoLocal(t *testing.T) {
	orig := Default()
	defer SetDefault(orig)
	SetDefault(Make("ja"))

	if Get().String() != "ja" {
		t.Errorf("Get() without local = %q, want ja", Get().String())
	}
}

func TestSet_Get(t *testing.T) {
	Set(Make("zh-TW"))
	defer Del()

	if Get().String() != "zh-TW" {
		t.Errorf("Get() = %q, want zh-TW", Get().String())
	}
}

func TestDelete(t *testing.T) {
	Set(Make("zh-TW"))
	Del()

	orig := Default()
	defer SetDefault(orig)
	SetDefault(Make("fr"))

	if Get().String() != "fr" {
		t.Errorf("Get() after Del() = %q, want fr (default)", Get().String())
	}
}

func TestGet_GoroutineIsolation(t *testing.T) {
	orig := Default()
	defer SetDefault(orig)
	SetDefault(Make("en"))

	var wg sync.WaitGroup
	wg.Add(3)

	results := make([]string, 3)

	// goroutine 0: set zh-CN
	go func() {
		defer wg.Done()
		Set(Make("zh-CN"))
		results[0] = Get().String()
	}()

	// goroutine 1: set ja
	go func() {
		defer wg.Done()
		Set(Make("ja"))
		results[1] = Get().String()
	}()

	// goroutine 2: no local override
	go func() {
		defer wg.Done()
		results[2] = Get().String()
	}()

	wg.Wait()

	if results[0] != "zh-CN" {
		t.Errorf("goroutine 0: got %q, want zh-CN", results[0])
	}
	if results[1] != "ja" {
		t.Errorf("goroutine 1: got %q, want ja", results[1])
	}
	if results[2] != "en" {
		t.Errorf("goroutine 2: got %q, want en (default)", results[2])
	}
}
