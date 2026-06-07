//go:build (lang_fr || lang_all) && (country_all || country_europe || country_im || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsleOfMan.RegisterName(xlanguage.French, "Île de Man")
	dataIsleOfMan.RegisterOfficialName(xlanguage.French, "Île de Man")
	dataIsleOfMan.RegisterCapital(xlanguage.French, "Douglas")
}
