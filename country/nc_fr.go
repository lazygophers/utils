//go:build country_all || country_melanesia || country_nc || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewCaledonia.RegisterName(xlanguage.French, "Nouvelle-Calédonie")
	dataNewCaledonia.RegisterOfficialName(xlanguage.French, "Nouvelle-Calédonie")
	dataNewCaledonia.RegisterCapital(xlanguage.French, "Nouméa")
}
