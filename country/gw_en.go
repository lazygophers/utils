//go:build country_africa || country_all || country_gw || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuineaBissau.RegisterName(xlanguage.English, "Guinea-Bissau")
	dataGuineaBissau.RegisterOfficialName(xlanguage.English, "Republic of Guinea-Bissau")
	dataGuineaBissau.RegisterCapital(xlanguage.English, "Bissau")
}
