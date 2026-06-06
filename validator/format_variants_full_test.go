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
	assert.Contains(t, formatMessageBuilder(allPH, withVal), "Name")
	assert.Contains(t, formatMessageBuilder(allPH, noVal), "Name")
	assert.Contains(t, formatMessageBuilder("{field}", noVal), "Name")
	assert.Contains(t, formatMessageBuilder("{param}", noVal), "5")
	assert.Contains(t, formatMessageBuilder("{tag}", noVal), "required")
	assert.Equal(t, "x", formatMessageBuilder("x", withVal))
}

func TestFmtByteSliceAll(t *testing.T) {
	assert.Contains(t, formatMessageByteSlice(allPH, withVal), "Name")
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
	assert.Contains(t, formatMessageNoFmt(allPH, withVal), "Name")
	assert.Contains(t, formatMessageNoFmt(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessageNoFmt("x", withVal))
}

func TestFmtSingleReplaceAll(t *testing.T) {
	assert.Contains(t, formatMessageSingleReplace(allPH, withVal), "Name")
	assert.Contains(t, formatMessageSingleReplace(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessageSingleReplace("x", withVal))
}

func TestFmtHashtableAll(t *testing.T) {
	assert.Contains(t, formatMessageHashtable(allPH, withVal), "Name")
	assert.Contains(t, formatMessageHashtable(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessageHashtable("x", withVal))
}

func TestFmtInlineCheckAll(t *testing.T) {
	assert.Contains(t, formatMessageInlineCheck(allPH, withVal), "Name")
	assert.Contains(t, formatMessageInlineCheck(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessageInlineCheck("x", withVal))
}

func TestFmtPrecomputeAll(t *testing.T) {
	assert.Contains(t, formatMessagePrecompute(allPH, withVal), "Name")
	assert.Contains(t, formatMessagePrecompute(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessagePrecompute("x", withVal))
}

func TestFmtBytesBufferAll(t *testing.T) {
	assert.Contains(t, formatMessageBytesBuffer(allPH, withVal), "Name")
	assert.Contains(t, formatMessageBytesBuffer(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessageBytesBuffer("x", withVal))
}

func TestFmtOptimizedCurrentAll(t *testing.T) {
	assert.Contains(t, formatMessageOptimizedCurrent(allPH, withVal), "Name")
	assert.Contains(t, formatMessageOptimizedCurrent(allPH, noVal), "Name")
	assert.Equal(t, "x", formatMessageOptimizedCurrent("x", withVal))
}

func TestFmtFastPathAll(t *testing.T) {
	assert.Contains(t, formatMessageFastPath(allPH, withVal), "Name")
	assert.Contains(t, formatMessageFastPath(allPH, noVal), "Name")
	assert.Contains(t, formatMessageFastPath("{field} {param} {tag}", noVal), "Name")
	assert.Equal(t, "x", formatMessageFastPath("x", withVal))
}

func TestCompileTemplateDirect(t *testing.T) {
	ct := compileTemplate("{field} x {tag}")
	assert.NotNil(t, ct)
	assert.True(t, len(ct.parts) >= 3)
}

func TestFormatCompiledNoPH(t *testing.T) {
	assert.Equal(t, "abc", formatMessageCompiled("abc", &FieldError{Field: "x"}))
}
