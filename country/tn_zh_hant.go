//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTunisia.RegisterName(xlanguage.MustParse("zh-Hant"), "突尼西亞")
	dataTunisia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "突尼西亞共和國")
	dataTunisia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "突尼斯")
}
