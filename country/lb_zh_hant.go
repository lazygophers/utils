//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_lb || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLebanon.RegisterName(xlanguage.MustParse("zh-Hant"), "黎巴嫩")
	dataLebanon.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "黎巴嫩共和國")
	dataLebanon.RegisterCapital(xlanguage.MustParse("zh-Hant"), "貝魯特")
}
