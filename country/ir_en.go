//go:build country_all || country_asia || country_ir || country_southern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIran.RegisterName(xlanguage.English, "Iran")
	dataIran.RegisterOfficialName(xlanguage.English, "Islamic Republic of Iran")
	dataIran.RegisterCapital(xlanguage.English, "Tehran")
}
