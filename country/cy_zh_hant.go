//go:build (lang_zh_hant || lang_all) && (country_all || country_cy || country_europe || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCyprus.RegisterName(xlanguage.MustParse("zh-Hant"), "賽普勒斯")
	dataCyprus.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "賽普勒斯共和國")
	dataCyprus.RegisterCapital(xlanguage.MustParse("zh-Hant"), "尼古西亞")
}
