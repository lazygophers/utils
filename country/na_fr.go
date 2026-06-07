//go:build (lang_fr || lang_all) && (country_africa || country_all || country_na || country_southern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNamibia.RegisterName(xlanguage.French, "Namibie")
	dataNamibia.RegisterOfficialName(xlanguage.French, "République de Namibie")
	dataNamibia.RegisterCapital(xlanguage.French, "Windhoek")
}
