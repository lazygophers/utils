package anyx_test

import (
	"github.com/lazygophers/utils/anyx"
	"testing"
)

type DeepStruct struct {
	Name string
	Age  int

	Childrens []*DeepStruct
	Map       map[string]*DeepStruct

	Deep *DeepStruct
}

func TestDeepCopy(t *testing.T) {
	a := &DeepStruct{
		Name: "a",
		Age:  1,
		Childrens: []*DeepStruct{
			{
				Name: "a1",
			},
		},
		Map: map[string]*DeepStruct{
			"a1": {
				Name: "a1",
			},
		},
		Deep: &DeepStruct{
			Name: "a2",
		},
	}

	b := &DeepStruct{
		Map: map[string]*DeepStruct{
			"a2": {
				Name: "a2",
			},
		},
	}

	anyx.DeepCopy(a, b)

	t.Logf("%+v", a)
	t.Logf("%+v", b)
}
