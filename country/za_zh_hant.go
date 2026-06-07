//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_southern_africa || country_za)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthAfrica.RegisterName(xlanguage.MustParse("zh-Hant"), "南非")
	dataSouthAfrica.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "南非共和國")
	dataSouthAfrica.RegisterCapital(xlanguage.MustParse("zh-Hant"), "普利托利亞")
}
