//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_ci || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIvoryCoast.RegisterName(xlanguage.MustParse("zh-Hant"), "象牙海岸")
	dataIvoryCoast.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "象牙海岸共和國")
	dataIvoryCoast.RegisterCapital(xlanguage.MustParse("zh-Hant"), "雅穆索戈")
}
