//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArgentina.RegisterName(xlanguage.MustParse("zh-Hant"), "阿根廷")
	dataArgentina.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "阿根廷共和國")
	dataArgentina.RegisterCapital(xlanguage.MustParse("zh-Hant"), "布宜諾斯艾利斯")
}
