//go:build country_all || country_au || country_australia_and_new_zealand || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustralia.RegisterName(xlanguage.English, "Australia")
	dataAustralia.RegisterOfficialName(xlanguage.English, "Commonwealth of Australia")
	dataAustralia.RegisterCapital(xlanguage.English, "Canberra")
}
