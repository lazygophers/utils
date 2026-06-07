//go:build (lang_ru || lang_all) && (country_all || country_micronesia || country_oceania || country_um)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsMinorOutlyingIslands.RegisterName(xlanguage.Russian, "Внешние малые острова США")
	dataUsMinorOutlyingIslands.RegisterOfficialName(xlanguage.Russian, "Внешние малые острова США")
}
