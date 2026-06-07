//go:build (lang_es || lang_all) && (country_all || country_asia || country_lb || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLebanon.RegisterName(xlanguage.Spanish, "Líbano")
	dataLebanon.RegisterOfficialName(xlanguage.Spanish, "República Libanesa")
	dataLebanon.RegisterCapital(xlanguage.Spanish, "Beirut")
}
