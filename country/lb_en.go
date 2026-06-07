//go:build country_all || country_asia || country_lb || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLebanon.RegisterName(xlanguage.English, "Lebanon")
	dataLebanon.RegisterOfficialName(xlanguage.English, "Lebanese Republic")
	dataLebanon.RegisterCapital(xlanguage.English, "Beirut")
}
