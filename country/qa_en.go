//go:build country_all || country_asia || country_qa || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataQatar.RegisterName(xlanguage.English, "Qatar")
	dataQatar.RegisterOfficialName(xlanguage.English, "State of Qatar")
	dataQatar.RegisterCapital(xlanguage.English, "Doha")
}
