//go:build (lang_fr || lang_all) && (country_africa || country_all || country_southern_africa || country_za)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthAfrica.RegisterName(xlanguage.French, "Afrique du Sud")
	dataSouthAfrica.RegisterOfficialName(xlanguage.French, "République d'Afrique du Sud")
	dataSouthAfrica.RegisterCapital(xlanguage.French, "Pretoria")
}
