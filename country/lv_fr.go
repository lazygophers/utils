//go:build (lang_fr || lang_all) && (country_all || country_europe || country_lv || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLatvia.RegisterName(xlanguage.French, "Lettonie")
	dataLatvia.RegisterOfficialName(xlanguage.French, "République de Lettonie")
	dataLatvia.RegisterCapital(xlanguage.French, "Riga")
}
