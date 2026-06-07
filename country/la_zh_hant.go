//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLaos.RegisterName(xlanguage.MustParse("zh-Hant"), "寮國")
	dataLaos.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "寮人民民主共和國")
	dataLaos.RegisterCapital(xlanguage.MustParse("zh-Hant"), "永珍")
}
