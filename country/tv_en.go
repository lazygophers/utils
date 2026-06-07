//go:build country_all || country_oceania || country_polynesia || country_tv

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTuvalu.RegisterName(xlanguage.English, "Tuvalu")
	dataTuvalu.RegisterOfficialName(xlanguage.English, "Tuvalu")
	dataTuvalu.RegisterCapital(xlanguage.English, "Funafuti")
}
