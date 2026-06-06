package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var allPH = "{field} must be {tag} with {param}, value={value}"
var noVal = &FieldError{Field: "Name", Tag: "required", Param: "5", Value: nil}
var withVal = &FieldError{Field: "Name", Tag: "min", Param: "10", Value: 42}

func TestFmtOrigAll(t *testing.T) {
	r := formatMessageOriginal(allPH, withVal)
	assert.Contains(t, r, "Name")
	assert.Contains(t, r, "42")
	r2 := formatMessageOriginal(allPH, noVal)
	assert.NotContains(t, r2, "42")
	assert.Equal(t, "x", formatMessageOriginal("x", withVal))
}

func TestFmtBuilderAll(t *testing.T) {
	r := formatMessageBuilder(allPH, withVal)
	assert.Contains(t, r, "Name")
	assert.Contains(t, r, "42", "{value} must be replaced")
	assert.Contains(t, formatMessageBuilder(allPH, noVal), "Name")
	assert.Contains(t, formatMessageBuilder("{field}", noVal), "Name")
	assert.Contains(t, formatMessageBuilder("{param}", noVal), "5")
	assert.Contains(t, formatMessageBuilder("{tag}", noVal), "required")
	assert.Contains(t, formatMessageBuilder("{value}", withVal), "42")
	assert.Equal(t, "x", formatMessageBuilder("x", withVal))
}

func TestFmtByteSliceAll(t *testing.T) {
	r := formatMessageByteSlice(allPH, withVal)
	assert.Contains(t, r, "Name")
	assert.Contains(t, r, "42", "{value} must be replaced")
	assert.Contains(t, formatMessageByteSlice(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessageByteSlice("x", withVal))
}

func TestFmtValueFastAll(t *testing.T) {
	assert.Equal(t, "hello", formatValueFast("hello"))
	assert.Equal(t, "42", formatValueFast(42))
	assert.Equal(t, "42", formatValueFast(int64(42)))
	assert.Contains(t, formatValueFast(3.14), "3.14")
	assert.Equal(t, "true", formatValueFast(true))
	assert.Equal(t, "false", formatValueFast(false))
	assert.Contains(t, formatValueFast([]int{1}), "[")
}

func TestFmtNoFmtAll(t *testing.T) {
	r := formatMessageNoFmt(allPH, withVal)
	assert.Contains(t, r, "Name")
	assert.Contains(t, r, "42", "{value} must be replaced")
	assert.Contains(t, formatMessageNoFmt(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessageNoFmt("x", withVal))
}

func TestFmtSingleReplaceAll(t *testing.T) {
	r := formatMessageSingleReplace(allPH, withVal)
	assert.Contains(t, r, "Name")
	assert.Contains(t, r, "42", "{value} must be replaced")
	assert.Contains(t, formatMessageSingleReplace(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessageSingleReplace("x", withVal))
}

func TestFmtHashtableAll(t *testing.T) {
	r := formatMessageHashtable(allPH, withVal)
	assert.Contains(t, r, "Name")
	assert.Contains(t, r, "42", "{value} must be replaced")
	assert.Contains(t, formatMessageHashtable(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessageHashtable("x", withVal))
}

func TestFmtInlineCheckAll(t *testing.T) {
	r := formatMessageInlineCheck(allPH, withVal)
	assert.Contains(t, r, "Name")
	assert.Contains(t, r, "42", "{value} must be replaced")
	assert.Contains(t, formatMessageInlineCheck(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessageInlineCheck("x", withVal))
}

func TestFmtPrecomputeAll(t *testing.T) {
	r := formatMessagePrecompute(allPH, withVal)
	assert.Contains(t, r, "Name")
	assert.Contains(t, r, "42", "{value} must be replaced")
	assert.Contains(t, formatMessagePrecompute(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessagePrecompute("x", withVal))
}

func TestFmtBytesBufferAll(t *testing.T) {
	r := formatMessageBytesBuffer(allPH, withVal)
	assert.Contains(t, r, "Name")
	assert.Contains(t, r, "42", "{value} must be replaced")
	assert.Contains(t, formatMessageBytesBuffer(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessageBytesBuffer("x", withVal))
}

func TestFmtOptimizedCurrentAll(t *testing.T) {
	r := formatMessageOptimizedCurrent(allPH, withVal)
	assert.Contains(t, r, "Name")
	assert.Contains(t, r, "42", "{value} must be replaced")
	assert.Contains(t, formatMessageOptimizedCurrent(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessageOptimizedCurrent("x", withVal))
}

func TestFmtFastPathAll(t *testing.T) {
	r := formatMessageFastPath(allPH, withVal)
	assert.Contains(t, r, "Name")
	assert.Contains(t, r, "42", "{value} must be replaced")
	assert.Contains(t, formatMessageFastPath(allPH, noVal), "Name")
	assert.Contains(t, formatMessageFastPath("{field} {param} {tag}", noVal), "Name")
	assert.Equal(t, "x", formatMessageFastPath("x", withVal))
}

func TestCompileTemplateDirect(t *testing.T) {
	ct := compileTemplate("{field} x {tag}")
	assert.NotNil(t, ct)
	assert.True(t, len(ct.parts) >= 3)
}

func TestCompileTemplateWithValue(t *testing.T) {
	ct := compileTemplate("{field}={value}")
	found := false
	for _, p := range ct.parts {
		if p.isPlaceholder && p.value == "value" {
			found = true
		}
	}
	assert.True(t, found, "compileTemplate must recognize {value} placeholder")
}

func TestFormatCompiledNoPH(t *testing.T) {
	assert.Equal(t, "abc", formatMessageCompiled("abc", &FieldError{Field: "x"}))
}

func TestFormatCompiledWithValue(t *testing.T) {
	r := formatMessageCompiled("{value}", withVal)
	assert.Equal(t, "42", r)
}

func TestFormatCompiledAllPH(t *testing.T) {
	r := formatMessageCompiled(allPH, withVal)
	assert.Contains(t, r, "Name")
	assert.Contains(t, r, "42")
	assert.Contains(t, r, "10")
	assert.Contains(t, r, "min")
}

func TestFormatCompiledNilValue(t *testing.T) {
	r := formatMessageCompiled("val={value}", noVal)
	assert.Equal(t, "val=", r)
}
