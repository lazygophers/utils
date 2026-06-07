//go:build (lang_fr || lang_all) && (country_all || country_cy || country_europe || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCyprus.RegisterName(xlanguage.French, "Chypre")
	dataCyprus.RegisterOfficialName(xlanguage.French, "République de Chypre")
	dataCyprus.RegisterCapital(xlanguage.French, "Nicosie")
}
