//go:build (lang_fr || lang_all) && (country_all || country_micronesia || country_mp || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthernMarianaIslands.RegisterName(xlanguage.French, "Îles Mariannes du Nord")
	dataNorthernMarianaIslands.RegisterOfficialName(xlanguage.French, "Commonwealth des Îles Mariannes du Nord")
	dataNorthernMarianaIslands.RegisterCapital(xlanguage.French, "Saipan")
}
