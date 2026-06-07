//go:build country_africa || country_all || country_southern_africa || country_sz

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEswatini.RegisterName(xlanguage.English, "Eswatini")
	dataEswatini.RegisterOfficialName(xlanguage.English, "Kingdom of Eswatini")
	dataEswatini.RegisterCapital(xlanguage.English, "Mbabane")
}
