//go:build (lang_fr || lang_all) && (country_all || country_oceania || country_polynesia || country_to)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTonga.RegisterName(xlanguage.French, "Tonga")
	dataTonga.RegisterOfficialName(xlanguage.French, "Royaume des Tonga")
	dataTonga.RegisterCapital(xlanguage.French, "Nukuʻalofa")
}
