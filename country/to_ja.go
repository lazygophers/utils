//go:build (lang_ja || lang_all) && (country_all || country_oceania || country_polynesia || country_to)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTonga.RegisterName(xlanguage.Japanese, "トンガ")
	dataTonga.RegisterOfficialName(xlanguage.Japanese, "トンガ王国")
	dataTonga.RegisterCapital(xlanguage.Japanese, "ヌクアロファ")
}
