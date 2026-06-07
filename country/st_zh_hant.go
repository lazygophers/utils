//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_middle_africa || country_st)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaoTomeAndPrincipe.RegisterName(xlanguage.MustParse("zh-Hant"), "聖多美普林西比")
	dataSaoTomeAndPrincipe.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "聖多美普林西比民主共和國")
	dataSaoTomeAndPrincipe.RegisterCapital(xlanguage.MustParse("zh-Hant"), "聖多美")
}
