package defaults_test

import (
	"github.com/lazygophers/utils/defaults"
	"testing"
)

type A struct {
	Name string `default:"name"`
}

func TestStruct(t *testing.T) {
	var s A
	err := defaults.SetDefaults(&s)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}

	t.Log(s)
}
