//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAndorra.RegisterName(xlanguage.MustParse("zh-Hant"), "安道爾")
	dataAndorra.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "安道爾公國")
	dataAndorra.RegisterCapital(xlanguage.MustParse("zh-Hant"), "老安道爾")
}
