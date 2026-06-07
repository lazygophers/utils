//go:build (lang_ja || lang_all) && (country_all || country_ch || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSwitzerland.RegisterName(xlanguage.Japanese, "スイス")
	dataSwitzerland.RegisterOfficialName(xlanguage.Japanese, "スイス連邦")
	dataSwitzerland.RegisterCapital(xlanguage.Japanese, "ベルン")
}
