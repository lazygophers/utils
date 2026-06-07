//go:build (lang_fr || lang_all) && (country_all || country_mh || country_micronesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMarshallIslands.RegisterName(xlanguage.French, "Îles Marshall")
	dataMarshallIslands.RegisterOfficialName(xlanguage.French, "République des Îles Marshall")
	dataMarshallIslands.RegisterCapital(xlanguage.French, "Majuro")
}
