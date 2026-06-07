//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_zm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZambia.RegisterName(xlanguage.Japanese, "ザンビア")
	dataZambia.RegisterOfficialName(xlanguage.Japanese, "ザンビア共和国")
	dataZambia.RegisterCapital(xlanguage.Japanese, "ルサカ")
}
