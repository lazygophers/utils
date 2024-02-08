package unit_test

import (
	"github.com/lazygophers/utils/unit"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	t.Log(unit.DurationMinuteSecond(time.Second * 10))
}
