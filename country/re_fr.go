//go:build country_africa || country_all || country_eastern_africa || country_re

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataReunion.RegisterName(xlanguage.French, "La Réunion")
	dataReunion.RegisterOfficialName(xlanguage.French, "La Réunion")
	dataReunion.RegisterCapital(xlanguage.French, "Saint-Denis")
}
