//go:build country_all || country_mh || country_micronesia || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMarshallIslands.RegisterName(xlanguage.English, "Marshall Islands")
	dataMarshallIslands.RegisterOfficialName(xlanguage.English, "Republic of the Marshall Islands")
	dataMarshallIslands.RegisterCapital(xlanguage.English, "Majuro")
}
