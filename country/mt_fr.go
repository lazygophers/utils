//go:build (lang_fr || lang_all) && (country_all || country_europe || country_mt || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalta.RegisterName(xlanguage.French, "Malte")
	dataMalta.RegisterOfficialName(xlanguage.French, "République de Malte")
	dataMalta.RegisterCapital(xlanguage.French, "La Valette")
}
