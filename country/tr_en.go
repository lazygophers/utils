//go:build country_all || country_asia || country_tr || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkey.RegisterName(xlanguage.English, "Turkey")
	dataTurkey.RegisterOfficialName(xlanguage.English, "Republic of Turkiye")
	dataTurkey.RegisterCapital(xlanguage.English, "Ankara")
}
