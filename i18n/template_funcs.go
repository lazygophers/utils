package i18n

import (
	"strings"
	"text/template"
	"time"

	"github.com/lazygophers/utils/candy"
	"github.com/lazygophers/utils/stringx"
)

// builtinTemplateFuncs 是 New 默认装载的模板函数表，覆盖常见字符串 / 切片 / 时间渲染需求。
// 用户可通过 WithTemplateFuncs / AddTemplateFunc 追加或覆盖同名函数。
var builtinTemplateFuncs = template.FuncMap{
	// strings 命名风格转换
	"ToCamel":      stringx.ToCamel,
	"ToSmallCamel": stringx.ToSmallCamel,
	"ToSnake":      stringx.ToSnake,
	"ToKebab":      stringx.ToKebab,
	"ToSlash":      stringx.ToSlash,
	"ToDot":        stringx.ToDot,

	// strings stdlib
	"ToLower":    strings.ToLower,
	"ToUpper":    strings.ToUpper,
	"ToTitle":    strings.ToTitle,
	"TrimPrefix": strings.TrimPrefix,
	"TrimSuffix": strings.TrimSuffix,
	"TrimSpace":  strings.TrimSpace,
	"HasPrefix":  strings.HasPrefix,
	"HasSuffix":  strings.HasSuffix,
	"Contains":   strings.Contains,
	"Replace":    strings.ReplaceAll,
	"Split":      strings.Split,
	"Join":       strings.Join,

	// candy 泛型切片便捷
	"UniqueString":   candy.Unique[string],
	"SortString":     candy.Sort[string],
	"ReverseString":  candy.Reverse[string],
	"FirstString":    candy.First[string],
	"LastString":     candy.Last[string],
	"ContainsString": candy.Contains[string],

	// 时间格式化
	"TimeFormat": func(t time.Time, layout string) string {
		return t.Format(layout)
	},
	"TimeFormatUnix": func(sec int64, layout string) string {
		return time.Unix(sec, 0).Format(layout)
	},
}
