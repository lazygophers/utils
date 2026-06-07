//go:build country_all || country_asia || country_iq || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIraq.RegisterName(xlanguage.English, "Iraq")
	dataIraq.RegisterOfficialName(xlanguage.English, "Republic of Iraq")
	dataIraq.RegisterCapital(xlanguage.English, "Baghdad")
}
