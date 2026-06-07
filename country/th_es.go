//go:build (lang_es || lang_all) && (country_all || country_asia || country_south_eastern_asia || country_th)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataThailand.RegisterName(xlanguage.Spanish, "Tailandia")
	dataThailand.RegisterOfficialName(xlanguage.Spanish, "Reino de Tailandia")
	dataThailand.RegisterCapital(xlanguage.Spanish, "Bangkok")
}
