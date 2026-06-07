//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGabon.RegisterName(xlanguage.MustParse("zh-Hant"), "加彭")
	dataGabon.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "加彭共和國")
	dataGabon.RegisterCapital(xlanguage.MustParse("zh-Hant"), "自由市")
}
