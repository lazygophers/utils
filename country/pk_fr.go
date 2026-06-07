//go:build (lang_fr || lang_all) && (country_all || country_asia || country_pk || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPakistan.RegisterName(xlanguage.French, "Pakistan")
	dataPakistan.RegisterOfficialName(xlanguage.French, "République islamique du Pakistan")
	dataPakistan.RegisterCapital(xlanguage.French, "Islamabad")
}
