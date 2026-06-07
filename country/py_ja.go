//go:build (lang_ja || lang_all) && (country_all || country_americas || country_py || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataParaguay.RegisterName(xlanguage.Japanese, "パラグアイ")
	dataParaguay.RegisterOfficialName(xlanguage.Japanese, "パラグアイ共和国")
	dataParaguay.RegisterCapital(xlanguage.Japanese, "アスンシオン")
}
