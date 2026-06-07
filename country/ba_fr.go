//go:build (lang_fr || lang_all) && (country_all || country_ba || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBosniaAndHerzegovina.RegisterName(xlanguage.French, "Bosnie-Herzégovine")
	dataBosniaAndHerzegovina.RegisterOfficialName(xlanguage.French, "Bosnie-Herzégovine")
	dataBosniaAndHerzegovina.RegisterCapital(xlanguage.French, "Sarajevo")
}
