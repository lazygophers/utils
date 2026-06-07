//go:build country_all || country_asia || country_mm || country_south_eastern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMyanmar.RegisterName(xlanguage.English, "Myanmar")
	dataMyanmar.RegisterOfficialName(xlanguage.English, "Republic of the Union of Myanmar")
	dataMyanmar.RegisterCapital(xlanguage.English, "Naypyidaw")
}
