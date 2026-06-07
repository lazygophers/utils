//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_sm || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSanMarino.RegisterName(xlanguage.MustParse("zh-Hant"), "聖馬利諾")
	dataSanMarino.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "聖馬利諾共和國")
	dataSanMarino.RegisterCapital(xlanguage.MustParse("zh-Hant"), "聖馬利諾")
}
