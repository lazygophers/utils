//go:build (lang_fr || lang_all) && (country_all || country_asia || country_id || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndonesia.RegisterName(xlanguage.French, "Indonésie")
	dataIndonesia.RegisterOfficialName(xlanguage.French, "République d'Indonésie")
	dataIndonesia.RegisterCapital(xlanguage.French, "Jakarta")
}
