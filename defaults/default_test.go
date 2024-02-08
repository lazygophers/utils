package defaults_test

import (
	"github.com/lazygophers/utils/defaults"
	"testing"
)

type TestStruct struct {
	Name string `default:"name"`
}

func TestSetDefault(t *testing.T) {
	var s TestStruct
	err := defaults.SetDefaults(&s)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}

	t.Log(s)
}
