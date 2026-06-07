//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_south_america || country_sr)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSuriname.RegisterName(xlanguage.MustParse("zh-Hant"), "蘇利南")
	dataSuriname.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "蘇利南共和國")
	dataSuriname.RegisterCapital(xlanguage.MustParse("zh-Hant"), "巴拉馬利波")
}
