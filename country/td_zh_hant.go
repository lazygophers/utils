//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChad.RegisterName(xlanguage.MustParse("zh-Hant"), "查德")
	dataChad.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "查德共和國")
	dataChad.RegisterCapital(xlanguage.MustParse("zh-Hant"), "恩將納")
}
