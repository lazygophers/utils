//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_hr || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCroatia.RegisterName(xlanguage.MustParse("zh-Hant"), "克羅埃西亞")
	dataCroatia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "克羅埃西亞共和國")
	dataCroatia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "札格瑞布")
}
