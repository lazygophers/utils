//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eastern_africa || country_tz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTanzania.RegisterName(xlanguage.French, "Tanzanie")
	dataTanzania.RegisterOfficialName(xlanguage.French, "République unie de Tanzanie")
	dataTanzania.RegisterCapital(xlanguage.French, "Dodoma")
}
