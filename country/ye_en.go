//go:build country_all || country_asia || country_western_asia || country_ye

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataYemen.RegisterName(xlanguage.English, "Yemen")
	dataYemen.RegisterOfficialName(xlanguage.English, "Republic of Yemen")
	dataYemen.RegisterCapital(xlanguage.English, "Sana'a")
}
