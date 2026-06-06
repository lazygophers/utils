package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/lazygophers/utils/language"
	xlanguage "golang.org/x/text/language"
)

// TestEffectiveLocaleGoroutineLocal 验证 goroutine-local 语言被自动使用
func TestEffectiveLocaleGoroutineLocal(t *testing.T) {
	// 保存并恢复
	origDefault := language.Default()
	defer language.SetDefault(origDefault)

	// 设置全局默认为 en（确保基线）
	language.SetDefault(language.Make("en"))

	v, err := New()
	require.NoError(t, err)

	// 基线：无 goroutine-local → 使用全局默认 en
	assert.Equal(t, xlanguage.Make("en"), v.GetLocale())

	// 设置 goroutine-local 为 zh
	language.Set(language.Make("zh"))
	defer language.Del()

	assert.Equal(t, xlanguage.Make("zh"), v.GetLocale())

	// 清除 goroutine-local → 回退到全局默认 en
	language.Del()
	assert.Equal(t, xlanguage.Make("en"), v.GetLocale())
}

// TestEffectiveLocaleGlobalDefault 验证 language.SetDefault 影响验证器
func TestEffectiveLocaleGlobalDefault(t *testing.T) {
	origDefault := language.Default()
	defer language.SetDefault(origDefault)

	language.SetDefault(language.Make("zh"))
	defer language.SetDefault(origDefault)

	language.Del() // 确保无 goroutine-local

	v, err := New()
	require.NoError(t, err)

	assert.Equal(t, xlanguage.Make("zh"), v.GetLocale())
}

// TestEffectiveLocaleExplicitOverridesGoroutineLocal 显式 locale 优先于 goroutine-local
func TestEffectiveLocaleExplicitOverridesGoroutineLocal(t *testing.T) {
	origDefault := language.Default()
	defer language.SetDefault(origDefault)
	language.SetDefault(language.Make("en"))

	language.Set(language.Make("zh"))
	defer language.Del()

	v, err := New(WithLocale(xlanguage.Make("ja")))
	require.NoError(t, err)

	// 显式设置 ja 优先于 goroutine-local zh
	assert.Equal(t, xlanguage.Make("ja"), v.GetLocale())
}

// TestTranslateWithGoroutineLocalLanguage goroutine 设置 zh 后，错误消息自动使用中文翻译
func TestTranslateWithGoroutineLocalLanguage(t *testing.T) {
	origDefault := language.Default()
	defer language.SetDefault(origDefault)
	language.SetDefault(language.Make("en"))

	language.Set(language.Make("zh"))
	defer language.Del()

	v, err := New()
	require.NoError(t, err)

	type User struct {
		Name string `validate:"required" json:"name"`
	}

	err2 := v.Struct(User{Name: ""})
	require.Error(t, err2)
	// zh locale 的 required 消息包含 "不能为空"
	assert.Contains(t, err2.Error(), "不能为空")
}

// TestTranslateFallbackToDefaultLanguage 无 goroutine-local 时使用全局默认翻译
func TestTranslateFallbackToDefaultLanguage(t *testing.T) {
	origDefault := language.Default()
	defer language.SetDefault(origDefault)

	language.SetDefault(language.Make("zh"))
	defer language.SetDefault(origDefault)
	language.Del()

	v, err := New()
	require.NoError(t, err)

	type User struct {
		Name string `validate:"required" json:"name"`
	}

	err2 := v.Struct(User{Name: ""})
	require.Error(t, err2)
	assert.Contains(t, err2.Error(), "不能为空")
}

// TestSetLocaleOverridesGoroutineLocal 运行时 SetLocale 覆盖 goroutine-local
func TestSetLocaleOverridesGoroutineLocal(t *testing.T) {
	origDefault := language.Default()
	defer language.SetDefault(origDefault)
	language.SetDefault(language.Make("en"))

	language.Set(language.Make("zh"))
	defer language.Del()

	v, err := New()
	require.NoError(t, err)

	// 当前是 goroutine-local zh
	assert.Equal(t, xlanguage.Make("zh"), v.GetLocale())

	// 显式设置 ja
	v.SetLocale(xlanguage.Make("ja"))
	assert.Equal(t, xlanguage.Make("ja"), v.GetLocale())

	// 清除显式设置 → 回到 goroutine-local zh
	v.SetLocale(xlanguage.Tag{})
	assert.Equal(t, xlanguage.Make("zh"), v.GetLocale())
}
