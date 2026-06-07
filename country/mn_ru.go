//go:build (lang_ru || lang_all) && (country_all || country_asia || country_eastern_asia || country_mn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMongolia.RegisterName(xlanguage.Russian, "Монголия")
	dataMongolia.RegisterOfficialName(xlanguage.Russian, "Монголия")
	dataMongolia.RegisterCapital(xlanguage.Russian, "Улан-Батор")
}
