//go:build (lang_ja || lang_all) && (country_all || country_asia || country_central_asia || country_tj)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTajikistan.RegisterName(xlanguage.Japanese, "タジキスタン")
	dataTajikistan.RegisterOfficialName(xlanguage.Japanese, "タジキスタン共和国")
	dataTajikistan.RegisterCapital(xlanguage.Japanese, "ドゥシャンベ")
}
