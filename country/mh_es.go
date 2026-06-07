//go:build (lang_es || lang_all) && (country_all || country_mh || country_micronesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMarshallIslands.RegisterName(xlanguage.Spanish, "Islas Marshall")
	dataMarshallIslands.RegisterOfficialName(xlanguage.Spanish, "República de las Islas Marshall")
	dataMarshallIslands.RegisterCapital(xlanguage.Spanish, "Majuro")
}
