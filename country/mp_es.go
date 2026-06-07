//go:build (lang_es || lang_all) && (country_all || country_micronesia || country_mp || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthernMarianaIslands.RegisterName(xlanguage.Spanish, "Islas Marianas del Norte")
	dataNorthernMarianaIslands.RegisterOfficialName(xlanguage.Spanish, "Mancomunidad de las Islas Marianas del Norte")
	dataNorthernMarianaIslands.RegisterCapital(xlanguage.Spanish, "Saipán")
}
