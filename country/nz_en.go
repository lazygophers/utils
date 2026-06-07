//go:build country_all || country_australia_and_new_zealand || country_nz || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewZealand.RegisterName(xlanguage.English, "New Zealand")
	dataNewZealand.RegisterOfficialName(xlanguage.English, "New Zealand")
	dataNewZealand.RegisterCapital(xlanguage.English, "Wellington")
}
