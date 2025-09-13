package candy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		name   string
		input  interface{}
		output interface{}
	}{
		{
			name:   "empty int slice",
			input:  []int{},
			output: 0,
		},
		{
			name:   "single int",
			input:  []int{5},
			output: 5,
		},
		{
			name:   "multiple ints",
			input:  []int{1, 2, 3, 4, 5},
			output: 3,
		},
		{
			name:   "negative ints",
			input:  []int{-1, -2, -3},
			output: -2,
		},
		{
			name:   "mixed positive negative",
			input:  []int{-5, 0, 5},
			output: 0,
		},
		{
			name:   "empty float64 slice",
			input:  []float64{},
			output: 0.0,
		},
		{
			name:   "single float64",
			input:  []float64{3.5},
			output: 3.5,
		},
		{
			name:   "multiple float64",
			input:  []float64{1.0, 2.0, 3.0},
			output: 2.0,
		},
		{
			name:   "float64 with precision",
			input:  []float64{1.1, 2.2, 3.3},
			output: 2.2,
		},
		{
			name:   "empty int32 slice",
			input:  []int32{},
			output: int32(0),
		},
		{
			name:   "int32 values",
			input:  []int32{10, 20, 30},
			output: int32(20),
		},
		{
			name:   "empty int64 slice",
			input:  []int64{},
			output: int64(0),
		},
		{
			name:   "int64 values",
			input:  []int64{100, 200, 300},
			output: int64(200),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch input := tt.input.(type) {
			case []int:
				result := Average(input)
				assert.Equal(t, tt.output, result)
			case []float64:
				result := Average(input)
				assert.InDelta(t, tt.output, result, 1e-9)
			case []int32:
				result := Average(input)
				assert.Equal(t, tt.output, result)
			case []int64:
				result := Average(input)
				assert.Equal(t, tt.output, result)
			}
		})
	}
}
