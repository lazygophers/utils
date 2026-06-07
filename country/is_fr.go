//go:build (lang_fr || lang_all) && (country_all || country_europe || country_is || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIceland.RegisterName(xlanguage.French, "Islande")
	dataIceland.RegisterOfficialName(xlanguage.French, "Islande")
	dataIceland.RegisterCapital(xlanguage.French, "Reykjavik")
}
