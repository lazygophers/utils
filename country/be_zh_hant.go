//go:build (lang_zh_hant || lang_all) && (country_all || country_be || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelgium.RegisterName(xlanguage.MustParse("zh-Hant"), "比利時")
	dataBelgium.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "比利時王國")
	dataBelgium.RegisterCapital(xlanguage.MustParse("zh-Hant"), "布魯塞爾")
}
