//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAzerbaijan.RegisterName(xlanguage.MustParse("zh-Hant"), "亞塞拜然")
	dataAzerbaijan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "亞塞拜然共和國")
	dataAzerbaijan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "巴庫")
}
