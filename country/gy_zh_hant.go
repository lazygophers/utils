//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_gy || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuyana.RegisterName(xlanguage.MustParse("zh-Hant"), "蓋亞那")
	dataGuyana.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "蓋亞那合作共和國")
	dataGuyana.RegisterCapital(xlanguage.MustParse("zh-Hant"), "喬治城")
}
