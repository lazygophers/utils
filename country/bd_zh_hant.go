//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_bd || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBangladesh.RegisterName(xlanguage.MustParse("zh-Hant"), "孟加拉")
	dataBangladesh.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "孟加拉人民共和國")
	dataBangladesh.RegisterCapital(xlanguage.MustParse("zh-Hant"), "達卡")
}
