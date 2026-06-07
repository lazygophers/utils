//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_sc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSeychelles.RegisterName(xlanguage.MustParse("zh-Hant"), "塞席爾")
	dataSeychelles.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "塞席爾共和國")
	dataSeychelles.RegisterCapital(xlanguage.MustParse("zh-Hant"), "維多利亞")
}
