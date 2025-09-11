package candy

import (
	"math"
	"testing"
)

func TestToBool(t *testing.T) {
	cases := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		// bool
		{"bool true", true, true},
		{"bool false", false, false},

		// int
		{"int 0", 0, false},
		{"int 1", 1, true},
		{"int -1", -1, true},
		{"int8 0", int8(0), false},
		{"int16 10", int16(10), true},
		{"int32 -10", int32(-10), true},
		{"int64 100", int64(100), true},

		// uint
		{"uint 0", uint(0), false},
		{"uint 1", uint(1), true},
		{"uint8 0", uint8(0), false},
		{"uint16 10", uint16(10), true},
		{"uint32 10", uint32(10), true},
		{"uint64 100", uint64(100), true},

		// float
		{"float32 0.0", float32(0.0), false},
		{"float32 -0.0", float32(-0.0), false},
		{"float32 1.23", float32(1.23), true},
		{"float32 -4.56", float32(-4.56), true},
		{"float32 NaN", float32(math.NaN()), false},
		{"float32 +Inf", float32(math.Inf(1)), true},
		{"float32 -Inf", float32(math.Inf(-1)), true},
		{"float64 0.0", 0.0, false},
		{"float64 -0.0", -0.0, false},
		{"float64 1.23", 1.23, true},
		{"float64 -4.56", -4.56, true},
		{"float64 NaN", math.NaN(), false},
		{"float64 +Inf", math.Inf(1), true},
		{"float64 -Inf", math.Inf(-1), true},

		// string (true values)
		{"string true", "true", true},
		{"string TRUE", "TRUE", true},
		{"string True", "True", true},
		{"string tRuE", "tRuE", true},
		{"string 1", "1", true},
		{"string t", "t", true},
		{"string T", "T", true},
		{"string y", "y", true},
		{"string Y", "Y", true},
		{"string yes", "yes", true},
		{"string YES", "YES", true},
		{"string Yes", "Yes", true},
		{"string yEs", "yEs", true},
		{"string on", "on", true},
		{"string ON", "ON", true},
		{"string On", "On", true},
		{"string oN", "oN", true},

		// string (false values)
		{"string false", "false", false},
		{"string FALSE", "FALSE", false},
		{"string False", "False", false},
		{"string fAlSe", "fAlSe", false},
		{"string 0", "0", false},
		{"string f", "f", false},
		{"string F", "F", false},
		{"string n", "n", false},
		{"string N", "N", false},
		{"string no", "no", false},
		{"string NO", "NO", false},
		{"string No", "No", false},
		{"string nO", "nO", false},
		{"string off", "off", false},
		{"string OFF", "OFF", false},
		{"string Off", "Off", false},
		{"string oFf", "oFf", false},

		// string (other non-empty)
		{"string hello", "hello", true},
		{"string with spaces", "  hello  ", true},

		// string (empty)
		{"string empty", "", false},
		{"string space only", "   ", false},
		{"string tab and newline", " \t\n \r\f\v ", false},

		// []byte (true values)
		{"[]byte true", []byte("true"), true},
		{"[]byte TRUE", []byte("TRUE"), true},
		{"[]byte True", []byte("True"), true},
		{"[]byte 1", []byte("1"), true},
		{"[]byte t", []byte("t"), true},
		{"[]byte T", []byte("T"), true},
		{"[]byte y", []byte("y"), true},
		{"[]byte Y", []byte("Y"), true},
		{"[]byte yes", []byte("yes"), true},
		{"[]byte YES", []byte("YES"), true},
		{"[]byte on", []byte("on"), true},
		{"[]byte ON", []byte("ON"), true},

		// []byte (false values)
		{"[]byte false", []byte("false"), false},
		{"[]byte FALSE", []byte("FALSE"), false},
		{"[]byte False", []byte("False"), false},
		{"[]byte 0", []byte("0"), false},
		{"[]byte f", []byte("f"), false},
		{"[]byte F", []byte("F"), false},
		{"[]byte n", []byte("n"), false},
		{"[]byte N", []byte("N"), false},
		{"[]byte no", []byte("no"), false},
		{"[]byte NO", []byte("NO"), false},
		{"[]byte off", []byte("off"), false},
		{"[]byte OFF", []byte("OFF"), false},

		// []byte (other non-empty)
		{"[]byte hello", []byte("hello"), true},
		{"[]byte with spaces", []byte("  hello  "), true},

		// []byte (empty)
		{"[]byte empty", []byte(""), false},
		{"[]byte space only", []byte("   "), false},
		{"[]byte tab and newline", []byte(" \t\n "), false},
		{"[]byte nil", []byte(nil), false},

		// nil
		{"nil", nil, false},

		// unsupported types
		{"unsupported struct", struct{}{}, false},
		{"unsupported map", make(map[int]int), false},
		{"unsupported slice", []int{1, 2}, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToBool(tc.input); got != tc.expected {
				t.Errorf("ToBool(%v) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}

func BenchmarkToBool(b *testing.B) {
	cases := []struct {
		input interface{}
	}{
		{true},
		{false},
		{1},
		{0},
		{1.23},
		{0.0},
		{float32(4.56)},
		{math.NaN()},
		{"true"},
		{"false"},
		{"1"},
		{"0"},
		{"t"},
		{"f"},
		{"hello"},
		{""},
		{[]byte("true")},
		{[]byte("false")},
		{[]byte("1")},
		{[]byte("0")},
		{[]byte("t")},
		{[]byte("f")},
		{[]byte("hello")},
		{[]byte("")},
		{[]byte("   ")},
		{nil},
		{struct{}{}},
	}

	b.ReportAllocs()
	b.ResetTimer()

	var r bool
	for i := 0; i < b.N; i++ {
		r = ToBool(cases[i%len(cases)].input)
	}
	_ = r
}