package unit_test

import (
	"github.com/lazygophers/utils/unit"
	"testing"
)

func TestSize(t *testing.T) {
	t.Log(unit.FormatSize(53687091200))
}

func TestSpeed(t *testing.T) {
	t.Log(unit.FormatSpeed(53687091200))
	t.Log(unit.Format2bps(53687091200))
}
