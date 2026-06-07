//go:build country_africa || country_all || country_eastern_africa || country_re

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataReunion.RegisterName(xlanguage.English, "Reunion")
	dataReunion.RegisterOfficialName(xlanguage.English, "Reunion Island")
	dataReunion.RegisterCapital(xlanguage.English, "Saint-Denis")
}
