//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_gg || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuernsey.RegisterName(xlanguage.MustParse("zh-Hant"), "根西")
	dataGuernsey.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "根西行政區")
	dataGuernsey.RegisterCapital(xlanguage.MustParse("zh-Hant"), "聖彼得港")
}
