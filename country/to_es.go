//go:build (lang_es || lang_all) && (country_all || country_oceania || country_polynesia || country_to)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTonga.RegisterName(xlanguage.Spanish, "Tonga")
	dataTonga.RegisterOfficialName(xlanguage.Spanish, "Reino de Tonga")
	dataTonga.RegisterCapital(xlanguage.Spanish, "Nukualofa")
}
