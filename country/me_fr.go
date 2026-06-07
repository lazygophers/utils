//go:build (lang_fr || lang_all) && (country_all || country_europe || country_me || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontenegro.RegisterName(xlanguage.French, "Monténégro")
	dataMontenegro.RegisterOfficialName(xlanguage.French, "Monténégro")
	dataMontenegro.RegisterCapital(xlanguage.French, "Podgorica")
}
