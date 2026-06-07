//go:build country_all || country_fj || country_melanesia || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFiji.RegisterName(xlanguage.English, "Fiji")
	dataFiji.RegisterOfficialName(xlanguage.English, "Republic of Fiji")
	dataFiji.RegisterCapital(xlanguage.English, "Suva")
}
